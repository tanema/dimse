package pdu

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/tanema/dimse/src/encoding"
)

type (
	PDU interface {
		String() string
		WritePayload(*encoding.Writer) error
	}
	Reader struct {
		reader *encoding.Reader
	}
	SubItem    interface{ Write(*encoding.Writer) error }
	AAssociate struct {
		Type            Type
		ProtocolVersion uint16
		CalledAETitle   string
		CallingAETitle  string
		Items           []SubItem
	}
	UserInformationItem                 struct{ Items []SubItem }
	UserInformationMaximumLengthItem    struct{ MaximumLengthReceived uint32 }
	AsynchronousOperationsWindowSubItem struct {
		MaxOpsInvoked   uint16
		MaxOpsPerformed uint16
	}
	RoleSelectionSubItem struct {
		SOPClassUID string
		SCURole     uint8
		SCPRole     uint8
	}
	// Container for subitems that this package doesnt' support
	SubItemUnsupported struct {
		Type Type
		Data []byte
	}
	ImplementationClassUIDSubItem    struct{ Name string }
	ImplementationVersionNameSubItem struct{ Name string }
	ApplicationContextItem           struct{ Name string }
	AbstractSyntaxSubItem            struct{ Name string }
	TransferSyntaxSubItem            struct{ Name string }
	PresentationContextItem          struct {
		Type      ItemType
		ContextID uint8
		Result    PresentationContextResult
		Items     []SubItem
	}
	AReleaseRq   struct{}
	AReleaseRp   struct{}
	AAssociateRj struct {
		Result RejectResultType
		Source SourceType
		Reason RejectReasonType
	}
	AAbort struct {
		Source SourceType
		Reason AbortReasonType
	}
	PDataTf struct {
		Items []PresentationDataValueItem
	}
	PresentationDataValueItem struct {
		ContextID uint8
		Command   bool // Bit 7 (LSB): 1 means command 0 means data
		Last      bool // Bit 6: 1 means last fragment. 0 means not last fragment.
		Value     []byte
	}
)

const (
	// The app context for DICOM. The first item in the A-ASSOCIATE-RQ
	DICOMApplicationContextItemName        = "1.2.840.10008.3.1.1.1"
	ImplementationClassUID                 = "1.2.826.0.1.3680043.9.7133"
	ImplementationName                     = "GODICOM_1_1"
	CurrentProtocolVersion          uint16 = 1
	DefaultMaxPDUSize               uint32 = 4 << 20
)

func (v *UserInformationItem) Write(w *encoding.Writer) error {
	enc := encoding.NewWriter(binary.BigEndian)
	for _, s := range v.Items {
		if err := s.Write(enc); err != nil {
			return err
		}
	}
	return w.Write(ItemTypeUserInformation, encoding.Skip(1), uint16(enc.Len()), enc.Bytes())
}

func (v *UserInformationMaximumLengthItem) Write(w *encoding.Writer) error {
	return w.Write(
		ItemTypeUserInformationMaximumLength,
		encoding.Skip(1),
		uint16(4),
		v.MaximumLengthReceived,
	)
}

func (v *ImplementationClassUIDSubItem) Write(e *encoding.Writer) error {
	return encSubItemWithName(e, ItemTypeImplementationClassUID, v.Name)
}

func (v *AsynchronousOperationsWindowSubItem) Write(w *encoding.Writer) error {
	return w.Write(
		ItemTypeAsynchronousOperationsWindow,
		encoding.Skip(1),
		uint16(4),
		v.MaxOpsPerformed,
		v.MaxOpsInvoked,
	)
}

func (v *RoleSelectionSubItem) Write(w *encoding.Writer) error {
	return w.Write(
		ItemTypeRoleSelection,
		encoding.Skip(1),
		uint16(2+len(v.SOPClassUID)+1*2),
		uint16(len(v.SOPClassUID)),
		v.SCURole,
		v.SCPRole,
	)
}

func (item *SubItemUnsupported) Write(w *encoding.Writer) error {
	return w.Write(
		item.Type,
		encoding.Skip(1),
		uint16(len(item.Data)),
		item.Data,
	)
}

func encSubItemWithName(w *encoding.Writer, itemType ItemType, name string) error {
	return w.Write(
		itemType,
		encoding.Skip(1),
		uint16(len(name)),
		[]byte(name),
	)
}

func (v *ImplementationVersionNameSubItem) Write(e *encoding.Writer) error {
	return encSubItemWithName(e, ItemTypeImplementationVersionName, v.Name)
}

func (v *ApplicationContextItem) Write(e *encoding.Writer) error {
	return encSubItemWithName(e, ItemTypeApplicationContext, v.Name)
}
func (v *AbstractSyntaxSubItem) Write(e *encoding.Writer) error {
	return encSubItemWithName(e, ItemTypeAbstractSyntax, v.Name)
}
func (v *TransferSyntaxSubItem) Write(e *encoding.Writer) error {
	return encSubItemWithName(e, ItemTypeTransferSyntax, v.Name)
}

func (v *PresentationContextItem) Write(w *encoding.Writer) error {
	if v.Type != ItemTypePresentationContextRequest && v.Type != ItemTypePresentationContextResponse {
		panic(*v)
	}
	enc := encoding.NewWriter(binary.BigEndian)
	for _, s := range v.Items {
		if err := s.Write(enc); err != nil {
			return err
		}
	}
	return w.Write(
		v.Type,
		encoding.Skip(1),
		uint16(4+enc.Len()),
		v.ContextID,
		encoding.Skip(3),
		enc.Bytes(),
	)
}

func (v *PresentationDataValueItem) Write(w *encoding.Writer) error {
	var header byte
	if v.Command {
		header |= 0b01
	}
	if v.Last {
		header |= 0b10
	}
	return w.Write(uint32(2+len(v.Value)), v.ContextID, header, v.Value)
}

// EncodePDU serializes "pdu" into []byte.
func EncodePDU(pdu PDU) ([]byte, error) {
	var pduType Type
	switch n := pdu.(type) {
	case *AAssociate:
		pduType = n.Type
	case *AAssociateRj:
		pduType = TypeAAssociateRj
	case *PDataTf:
		pduType = TypePDataTf
	case *AReleaseRq:
		pduType = TypeAReleaseRq
	case *AReleaseRp:
		pduType = TypeAReleaseRp
	case *AAbort:
		pduType = TypeAAbort
	default:
		panic(fmt.Sprintf("Unknown PDU %v", pdu))
	}
	enc := encoding.NewWriter(binary.BigEndian)
	if err := pdu.WritePayload(enc); err != nil {
		return nil, err
	}

	var header [6]byte // First 6 bytes of buf.
	header[0] = uint8(pduType)
	header[1] = 0 // Reserved.
	binary.BigEndian.PutUint32(header[2:6], uint32(enc.Len()))
	return append(header[:], enc.Bytes()...), nil
}

func NewReader(in io.Reader) *Reader {
	return &Reader{
		reader: encoding.NewReader(in, binary.BigEndian),
	}
}

func (r *Reader) Next() (PDU, error) {
	var pduType Type
	var length uint32
	if err := r.reader.Read(&pduType, encoding.Skip(1), &length); err != nil {
		return nil, fmt.Errorf("ReadPDUs error reading pdu header: %v", err)
	}
	return r.pdu(pduType, int(length))
}

// EncodePDU reads a "pdu" from a stream. maxPDUSize defines the maximum
// possible PDU size, in bytes, accepted by the caller.
func (r *Reader) pdu(pduType Type, length int) (PDU, error) {
	r.reader.PushLimit(int(length))
	switch pduType {
	case TypeAAssociateRq, TypeAAssociateAc:
		assoc := &AAssociate{}
		assoc.Type = pduType
		if err := r.reader.Read(&assoc.ProtocolVersion, encoding.Skip(2)); err != nil {
			return assoc, fmt.Errorf("error reading protocol version %v", err)
		} else if assoc.CalledAETitle, err = r.reader.String(16); err != nil {
			return assoc, fmt.Errorf("error reading called aetitle %v", err)
		} else if assoc.CallingAETitle, err = r.reader.String(16); err != nil {
			return assoc, fmt.Errorf("error reading calling aetitle %v", err)
		} else if err := r.reader.Read(encoding.Skip(8 * 4)); err != nil {
			return assoc, fmt.Errorf("error reading reserved data chunk %v", err)
		}
		for {
			item, err := r.decodeSubItem()
			if err != nil {
				if errors.Is(err, io.EOF) {
					return assoc, nil
				}
				return assoc, err
			}
			assoc.Items = append(assoc.Items, item)
		}
	case TypeAAssociateRj:
		if err := r.reader.Read(encoding.Skip(1)); err != nil { // reserved
			return nil, err
		}
		assoc := &AAssociateRj{}
		return assoc, r.reader.Read(&assoc.Result, &assoc.Source, &assoc.Reason)
	case TypeAAbort:
		if err := r.reader.Read(encoding.Skip(2)); err != nil {
			return nil, err
		}
		abort := &AAbort{}
		return abort, r.reader.Read(&abort.Source, &abort.Reason)
	case TypePDataTf:
		pdtf := &PDataTf{}
		for {
			item := PresentationDataValueItem{}
			var length uint32
			var header uint8
			if err := r.reader.Read(&length, &item.ContextID, &header); err != nil {
				if errors.Is(err, io.EOF) {
					return pdtf, nil
				}
				return pdtf, fmt.Errorf("error reading pdata header: %v", err)
			}
			item.Command = (header&1 != 0)
			item.Last = (header&2 != 0)
			item.Value = make([]byte, int(length-2))
			if err := r.reader.Read(&item.Value); err != nil {
				if errors.Is(err, io.EOF) {
					return pdtf, nil
				}
				return pdtf, fmt.Errorf("error reading pdata value: %v", err)
			}
			pdtf.Items = append(pdtf.Items, item)
		}
	case TypeAReleaseRq:
		return &AReleaseRq{}, r.reader.Read(encoding.Skip(4))
	case TypeAReleaseRp:
		return &AReleaseRp{}, r.reader.Read(encoding.Skip(4))
	default:
		return nil, fmt.Errorf("ReadPDU: unknown message type %d", pduType)
	}
}

func (r *Reader) decodeSubItem() (SubItem, error) {
	var itemType ItemType
	var length uint16
	if err := r.reader.Read(&itemType, encoding.Skip(1), &length); err != nil {
		return nil, err
	}

	switch itemType {
	case ItemTypeApplicationContext:
		name, err := r.reader.String(int(length))
		return &ApplicationContextItem{Name: name}, err
	case ItemTypeAbstractSyntax:
		name, err := r.reader.String(int(length))
		return &AbstractSyntaxSubItem{Name: name}, err
	case ItemTypeTransferSyntax:
		name, err := r.reader.String(int(length))
		return &TransferSyntaxSubItem{Name: name}, err
	case ItemTypeImplementationClassUID:
		name, err := r.reader.String(int(length))
		return &ImplementationClassUIDSubItem{Name: name}, err
	case ItemTypeImplementationVersionName:
		name, err := r.reader.String(int(length))
		return &ImplementationVersionNameSubItem{Name: name}, err
	case ItemTypeUserInformationMaximumLength:
		if length != 4 {
			return nil, fmt.Errorf("UserInformationMaximumLengthItem must be 4 bytes, but found %dB", length)
		}
		var maxLen uint32
		return &UserInformationMaximumLengthItem{MaximumLengthReceived: maxLen}, r.reader.Read(&maxLen)
	case ItemTypeAsynchronousOperationsWindow:
		var maxOpsInv, maxOpsPerf uint16
		if err := r.reader.Read(&maxOpsInv); err != nil {
			return nil, err
		}
		return &AsynchronousOperationsWindowSubItem{MaxOpsInvoked: maxOpsInv, MaxOpsPerformed: maxOpsPerf}, r.reader.Read(&maxOpsPerf)
	case ItemTypeRoleSelection:
		var uidLen uint16
		var scuRole, scpRole byte
		if err := r.reader.Read(&uidLen); err != nil {
			return nil, err
		}
		sopClassUID, err := r.reader.String(int(uidLen))
		if err != nil {
			return nil, err
		}
		return &RoleSelectionSubItem{SOPClassUID: sopClassUID, SCURole: scuRole, SCPRole: scpRole}, r.reader.Read(&scuRole, &scpRole)
	case ItemTypePresentationContextRequest, ItemTypePresentationContextResponse:
		v := &PresentationContextItem{Type: itemType}
		r.reader.PushLimit(int(length))
		defer r.reader.PopLimit()
		if err := r.reader.Read(&v.ContextID, encoding.Skip(1), &v.Result, encoding.Skip(1)); err != nil {
			return nil, err
		} else if v.ContextID%2 != 1 {
			return nil, fmt.Errorf("PresentationContextItem ID must be odd, but found %x", v.ContextID)
		}
		for {
			item, err := r.decodeSubItem()
			if err != nil {
				if errors.Is(err, io.EOF) {
					return v, nil
				}
				return v, err
			}
			v.Items = append(v.Items, item)
		}
	case ItemTypeUserInformation:
		v := &UserInformationItem{}
		r.reader.PushLimit(int(length))
		defer r.reader.PopLimit()
		for {
			item, err := r.decodeSubItem()
			if err != nil {
				if errors.Is(err, io.EOF) {
					return v, nil
				}
				return v, err
			}
			v.Items = append(v.Items, item)
		}
	default:
		return nil, fmt.Errorf("Unknown item type: 0x%x", itemType)
	}
}

func (pdu *AReleaseRq) WritePayload(w *encoding.Writer) error { return w.Write(encoding.Skip(4)) }
func (pdu *AReleaseRq) String() string                        { return fmt.Sprintf("Release Request") }

func (pdu *AReleaseRp) WritePayload(w *encoding.Writer) error { return w.Write(encoding.Skip(4)) }
func (pdu *AReleaseRp) String() string                        { return fmt.Sprintf("Release Response") }

func (pdu *AAssociate) WritePayload(w *encoding.Writer) error {
	if pdu.Type == 0 || pdu.CalledAETitle == "" || pdu.CallingAETitle == "" {
		return fmt.Errorf("Malformed associate")
	}
	if err := w.Write(
		pdu.ProtocolVersion,
		encoding.Skip(2),
		[]byte(padString(pdu.CalledAETitle, 16)),
		[]byte(padString(pdu.CallingAETitle, 16)),
		encoding.Skip(32),
	); err != nil {
		return err
	}

	for _, item := range pdu.Items {
		if err := item.Write(w); err != nil {
			return err
		}
	}
	return nil
}

func (pdu *AAssociate) String() string {
	return fmt.Sprintf("%s", pdu.Type)
}

func (pdu *AAssociateRj) WritePayload(w *encoding.Writer) error {
	return w.Write(encoding.Skip(1), &pdu.Result, &pdu.Source, &pdu.Reason)
}
func (pdu *AAssociateRj) String() string { return fmt.Sprintf("Associate Rejection") }

func (pdu *AAbort) WritePayload(w *encoding.Writer) error {
	return w.Write(encoding.Skip(2), &pdu.Source, &pdu.Reason)
}
func (pdu *AAbort) String() string { return fmt.Sprintf("Abort") }

func (pdu *PDataTf) WritePayload(w *encoding.Writer) error {
	for _, item := range pdu.Items {
		if err := item.Write(w); err != nil {
			return err
		}
	}
	return nil
}

func (pdu *PDataTf) String() string { return fmt.Sprintf("PData") }

// padString pads the string with " " up to the given length.
func padString(v string, length int) string {
	if len(v) > length {
		return v[:length]
	}
	return v + strings.Repeat(" ", length-len(v))
}
