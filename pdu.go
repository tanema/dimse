package dimse

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"strings"
)

type (
	PDU        interface{ WritePayload(*pduEncoder) error }
	SubItem    interface{ Write(*pduEncoder) error }
	AAssociate struct {
		Type            byte
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
		SCURole     byte
		SCPRole     byte
	}
	// Container for subitems that this package doesnt' support
	SubItemUnsupported struct {
		Type byte
		Data []byte
	}
	ImplementationClassUIDSubItem    struct{ Name string }
	ImplementationVersionNameSubItem struct{ Name string }
	ApplicationContextItem           struct{ Name string }
	AbstractSyntaxSubItem            struct{ Name string }
	TransferSyntaxSubItem            struct{ Name string }
	PresentationContextItem          struct {
		Type      byte
		ContextID byte
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
	PresentationContextResult byte
	RejectResultType          byte
	RejectReasonType          byte
	SourceType                byte
	AbortReasonType           byte
	AAbort                    struct {
		Source SourceType
		Reason AbortReasonType
	}
	PDataTf struct {
		Items []PresentationDataValueItem
	}
	PresentationDataValueItem struct {
		ContextID byte
		Command   bool // Bit 7 (LSB): 1 means command 0 means data
		Last      bool // Bit 6: 1 means last fragment. 0 means not last fragment.
		Value     []byte
	}
)

const (
	TypeAAssociateRq byte = 1 // A_ASSOCIATE_RQ association request
	TypeAAssociateAc byte = 2 // A_ASSOCIATE_AC association accepted
	TypeAAssociateRj byte = 3 // A_ASSOCIATE_RJ association rejuected
	TypePDataTf      byte = 4 // P_DATA_TF      used once an association has been established to send DIMSE message data.
	TypeAReleaseRq   byte = 5 // A_RELEASE_RQ   disassociation request
	TypeAReleaseRp   byte = 6 // A_RELEASE_RP   disassociation response
	TypeAAbort       byte = 7 // A_ABORT        disassociation without response.

	// Possible Type field values for SubItem.
	ItemTypeApplicationContext           byte = 0x10
	ItemTypePresentationContextRequest   byte = 0x20
	ItemTypePresentationContextResponse  byte = 0x21
	ItemTypeAbstractSyntax               byte = 0x30
	ItemTypeTransferSyntax               byte = 0x40
	ItemTypeUserInformation              byte = 0x50
	ItemTypeUserInformationMaximumLength byte = 0x51
	ItemTypeImplementationClassUID       byte = 0x52
	ItemTypeAsynchronousOperationsWindow byte = 0x53
	ItemTypeRoleSelection                byte = 0x54
	ItemTypeImplementationVersionName    byte = 0x55

	CurrentProtocolVersion uint16 = 1

	PresentationContextAccepted                                    PresentationContextResult = 0
	PresentationContextUserRejection                               PresentationContextResult = 1
	PresentationContextProviderRejectionNoReason                   PresentationContextResult = 2
	PresentationContextProviderRejectionAbstractSyntaxNotSupported PresentationContextResult = 3
	PresentationContextProviderRejectionTransferSyntaxNotSupported PresentationContextResult = 4

	ResultRejectedPermanent RejectResultType = 1
	ResultRejectedTransient RejectResultType = 2

	RejectReasonNone                               RejectReasonType = 1
	RejectReasonApplicationContextNameNotSupported RejectReasonType = 2
	RejectReasonCallingAETitleNotRecognized        RejectReasonType = 3
	RejectReasonCalledAETitleNotRecognized         RejectReasonType = 7

	SourceULServiceUser                 SourceType = 1
	SourceULServiceProviderACSE         SourceType = 2
	SourceULServiceProviderPresentation SourceType = 3

	AbortReasonNotSpecified             AbortReasonType = 0
	AbortReasonUnexpectedPDU            AbortReasonType = 2
	AbortReasonUnrecognizedPDUParameter AbortReasonType = 3
	AbortReasonUnexpectedPDUParameter   AbortReasonType = 4
	AbortReasonInvalidPDUParameterValue AbortReasonType = 5

	// The app context for DICOM. The first item in the A-ASSOCIATE-RQ
	DICOMApplicationContextItemName = "1.2.840.10008.3.1.1.1"
)

func (v *UserInformationItem) Write(w *pduEncoder) error {
	buf := bytes.NewBuffer(nil)
	enc := newEncoder(buf)
	for _, s := range v.Items {
		if err := s.Write(enc); err != nil {
			return err
		}
	}
	return w.Write(ItemTypeUserInformation, encSkip(1), uint16(buf.Len()), buf.Bytes())
}

func (v *UserInformationMaximumLengthItem) Write(w *pduEncoder) error {
	return w.Write(
		ItemTypeUserInformationMaximumLength,
		encSkip(1),
		uint16(4),
		v.MaximumLengthReceived,
	)
}

func (v *ImplementationClassUIDSubItem) Write(e *pduEncoder) error {
	return encSubItemWithName(e, ItemTypeImplementationClassUID, v.Name)
}

func (v *AsynchronousOperationsWindowSubItem) Write(w *pduEncoder) error {
	return w.Write(
		ItemTypeAsynchronousOperationsWindow,
		encSkip(1),
		uint16(4),
		v.MaxOpsPerformed,
		v.MaxOpsInvoked,
	)
}

func (v *RoleSelectionSubItem) Write(w *pduEncoder) error {
	return w.Write(
		ItemTypeRoleSelection,
		encSkip(1),
		uint16(2+len(v.SOPClassUID)+1*2),
		uint16(len(v.SOPClassUID)),
		v.SCURole,
		v.SCPRole,
	)
}

func (item *SubItemUnsupported) Write(w *pduEncoder) error {
	return w.Write(
		item.Type,
		encSkip(1),
		uint16(len(item.Data)),
		item.Data,
	)
}

func encSubItemWithName(w *pduEncoder, itemType byte, name string) error {
	return w.Write(
		itemType,
		encSkip(1),
		uint16(len(name)),
		[]byte(name),
	)
}

func (v *ImplementationVersionNameSubItem) Write(e *pduEncoder) error {
	return encSubItemWithName(e, ItemTypeImplementationVersionName, v.Name)
}

func (v *ApplicationContextItem) Write(e *pduEncoder) error {
	return encSubItemWithName(e, ItemTypeApplicationContext, v.Name)
}
func (v *AbstractSyntaxSubItem) Write(e *pduEncoder) error {
	return encSubItemWithName(e, ItemTypeAbstractSyntax, v.Name)
}
func (v *TransferSyntaxSubItem) Write(e *pduEncoder) error {
	return encSubItemWithName(e, ItemTypeTransferSyntax, v.Name)
}

func (v *PresentationContextItem) Write(w *pduEncoder) error {
	if v.Type != ItemTypePresentationContextRequest && v.Type != ItemTypePresentationContextResponse {
		panic(*v)
	}
	buf := bytes.NewBuffer(nil)
	enc := newEncoder(buf)
	for _, s := range v.Items {
		if err := s.Write(enc); err != nil {
			return err
		}
	}
	return w.Write(
		v.Type,
		encSkip(1),
		uint16(4+buf.Len()),
		v.ContextID,
		encSkip(3),
		buf.Bytes(),
	)
}

func (v *PresentationDataValueItem) Write(w *pduEncoder) error {
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
	var pduType byte
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
	buf := bytes.NewBuffer(nil)
	if err := pdu.WritePayload(newEncoder(buf)); err != nil {
		return nil, err
	}

	var header [6]byte // First 6 bytes of buf.
	header[0] = byte(pduType)
	header[1] = 0 // Reserved.
	binary.BigEndian.PutUint32(header[2:6], uint32(buf.Len()))
	return append(header[:], buf.Bytes()...), nil
}

// EncodePDU reads a "pdu" from a stream. maxPDUSize defines the maximum
// possible PDU size, in bytes, accepted by the caller.
func ReadPDU(in io.Reader) (PDU, error) {
	var pduType byte
	var length uint32
	d := newDecoder(in)

	if err := d.Read(&pduType, encSkip(1), &length); err != nil {
		return nil, err
	}

	d.PushLimit(int(length))
	switch pduType {
	case TypeAAssociateRq, TypeAAssociateAc:
		assoc := &AAssociate{}
		assoc.Type = pduType
		if err := d.Read(&assoc.ProtocolVersion, encSkip(2)); err != nil {
			return assoc, err
		} else if assoc.CalledAETitle, err = d.String(16); err != nil {
			return assoc, err
		} else if assoc.CallingAETitle, err = d.String(16); err != nil {
			return assoc, err
		} else if err := d.Read(encSkip(8 * 4)); err != nil {
			return assoc, err
		} else if assoc.CalledAETitle == "" || assoc.CallingAETitle == "" {
			return assoc, fmt.Errorf("A_ASSOCIATE.{Called,Calling}AETitle must not be empty")
		}
		for {
			item, err := decodeSubItem(d)
			if err != nil {
				if errors.Is(err, io.EOF) {
					return assoc, nil
				}
				return assoc, err
			}
			assoc.Items = append(assoc.Items, item)
		}
	case TypeAAssociateRj:
		if err := d.Read(encSkip(1)); err != nil { // reserved
			return nil, err
		}
		assoc := &AAssociateRj{}
		return assoc, d.Read(&assoc.Result, &assoc.Source, &assoc.Reason)
	case TypeAAbort:
		if err := d.Read(encSkip(2)); err != nil {
			return nil, err
		}
		abort := &AAbort{}
		return abort, d.Read(&abort.Source, &abort.Reason)
	case TypePDataTf:
		pdtf := &PDataTf{}
		for {
			item := PresentationDataValueItem{}
			var length uint32
			var header byte
			if err := d.Read(&length, &item.ContextID, &header); err != nil {
				if errors.Is(err, io.EOF) {
					return pdtf, nil
				}
				return pdtf, err
			}
			item.Command = (header&1 != 0)
			item.Last = (header&2 != 0)
			item.Value = make([]byte, int(length-2))
			if err := d.Read(&item.Value); err != nil {
				if errors.Is(err, io.EOF) {
					return pdtf, nil
				}
				return pdtf, err
			}
			pdtf.Items = append(pdtf.Items, item)
		}
	case TypeAReleaseRq:
		return &AReleaseRq{}, d.Read(encSkip(4))
	case TypeAReleaseRp:
		return &AReleaseRp{}, d.Read(encSkip(4))
	default:
		return nil, fmt.Errorf("ReadPDU: unknown message type %d", pduType)
	}
}

func decodeSubItem(d *pduDecoder) (SubItem, error) {
	var itemType byte
	var length uint16
	if err := d.Read(&itemType, encSkip(1), &length); err != nil {
		return nil, err
	}

	switch itemType {
	case ItemTypeApplicationContext:
		name, err := d.String(int(length))
		return &ApplicationContextItem{Name: name}, err
	case ItemTypeAbstractSyntax:
		name, err := d.String(int(length))
		return &AbstractSyntaxSubItem{Name: name}, err
	case ItemTypeTransferSyntax:
		name, err := d.String(int(length))
		return &TransferSyntaxSubItem{Name: name}, err
	case ItemTypeImplementationClassUID:
		name, err := d.String(int(length))
		return &ImplementationClassUIDSubItem{Name: name}, err
	case ItemTypeImplementationVersionName:
		name, err := d.String(int(length))
		return &ImplementationVersionNameSubItem{Name: name}, err
	case ItemTypeUserInformationMaximumLength:
		if length != 4 {
			return nil, fmt.Errorf("UserInformationMaximumLengthItem must be 4 bytes, but found %dB", length)
		}
		var maxLen uint32
		return &UserInformationMaximumLengthItem{MaximumLengthReceived: maxLen}, d.Read(&maxLen)
	case ItemTypeAsynchronousOperationsWindow:
		var maxOpsInv, maxOpsPerf uint16
		if err := d.Read(&maxOpsInv); err != nil {
			return nil, err
		}
		return &AsynchronousOperationsWindowSubItem{MaxOpsInvoked: maxOpsInv, MaxOpsPerformed: maxOpsPerf}, d.Read(&maxOpsPerf)
	case ItemTypeRoleSelection:
		var uidLen uint16
		var scuRole, scpRole byte
		if err := d.Read(&uidLen); err != nil {
			return nil, err
		}
		sopClassUID, err := d.String(int(uidLen))
		if err != nil {
			return nil, err
		}
		return &RoleSelectionSubItem{SOPClassUID: sopClassUID, SCURole: scuRole, SCPRole: scpRole}, d.Read(&scuRole, &scpRole)
	case ItemTypePresentationContextRequest, ItemTypePresentationContextResponse:
		v := &PresentationContextItem{Type: itemType}
		d.PushLimit(int(length))
		defer d.PopLimit()
		if err := d.Read(&v.ContextID, encSkip(1), &v.Result, encSkip(1)); err != nil {
			return nil, err
		} else if v.ContextID%2 != 1 {
			return nil, fmt.Errorf("PresentationContextItem ID must be odd, but found %x", v.ContextID)
		}
		for {
			item, err := decodeSubItem(d)
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
		d.PushLimit(int(length))
		defer d.PopLimit()
		for {
			item, err := decodeSubItem(d)
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

func (pdu *AReleaseRq) WritePayload(w *pduEncoder) error { return w.Write(encSkip(4)) }
func (pdu *AReleaseRp) WritePayload(w *pduEncoder) error { return w.Write(encSkip(4)) }

func (pdu *AAssociate) WritePayload(w *pduEncoder) error {
	if pdu.Type == 0 || pdu.CalledAETitle == "" || pdu.CallingAETitle == "" {
		return fmt.Errorf("Malformed associate")
	}
	if err := w.Write(
		pdu.ProtocolVersion,
		encSkip(2),
		[]byte(padString(pdu.CalledAETitle, 16)),
		[]byte(padString(pdu.CallingAETitle, 16)),
		encSkip(32),
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

func (pdu *AAssociateRj) WritePayload(w *pduEncoder) error {
	return w.Write(encSkip(1), &pdu.Result, &pdu.Source, &pdu.Reason)
}

func (pdu *AAbort) WritePayload(w *pduEncoder) error {
	return w.Write(encSkip(2), &pdu.Source, &pdu.Reason)
}

func (pdu *PDataTf) WritePayload(w *pduEncoder) error {
	for _, item := range pdu.Items {
		if err := item.Write(w); err != nil {
			return err
		}
	}
	return nil
}

// padString pads the string with " " up to the given length.
func padString(v string, length int) string {
	if len(v) > length {
		return v[:length]
	}
	return v + strings.Repeat(" ", length-len(v))
}
