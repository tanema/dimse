package pdu

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/tanema/dimse/src/encoding"
)

type Reader struct {
	reader *encoding.Reader
}

func NewReader(in io.Reader) *Reader {
	return &Reader{
		reader: encoding.NewReader(in, binary.BigEndian),
	}
}

func (r *Reader) Next() (any, error) {
	var pduType Type
	var length uint32
	if err := r.reader.Read(&pduType, encoding.Skip(1), &length); err != nil {
		return nil, fmt.Errorf("ReadPDUs error reading pdu header: %v", err)
	}
	return r.pdu(pduType, int(length))
}

// EncodePDU reads a "pdu" from a stream. maxPDUSize defines the maximum
// possible PDU size, in bytes, accepted by the caller.
func (r *Reader) pdu(pduType Type, length int) (any, error) {
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

func (r *Reader) decodeSubItem() (any, error) {
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
