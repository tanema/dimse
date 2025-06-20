package pdu

import (
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/tanema/dimse/src/encoding"
)

// EncodePDU serializes "pdu" into []byte.
func EncodePDU(pdu any) ([]byte, error) {
	enc := encoding.NewWriter(binary.BigEndian)
	var pduType Type
	var err error
	switch n := pdu.(type) {
	case *AAssociate:
		pduType = n.Type
		if n.Type == 0 || n.CalledAETitle == "" || n.CallingAETitle == "" {
			return nil, fmt.Errorf("Malformed associate")
		}
		if err := enc.Write(
			n.ProtocolVersion,
			encoding.Skip(2),
			[]byte(padString(n.CalledAETitle, 16)),
			[]byte(padString(n.CallingAETitle, 16)),
			encoding.Skip(32),
		); err != nil {
			return nil, err
		}

		for _, item := range n.Items {
			if err := writeSubItem(enc, item); err != nil {
				return nil, err
			}
		}
	case *AAssociateRj:
		pduType = TypeAAssociateRj
		err = enc.Write(encoding.Skip(1), &n.Result, &n.Source, &n.Reason)
	case *PDataTf:
		pduType = TypePDataTf
		for _, item := range n.Items {
			if err := writeSubItem(enc, item); err != nil {
				return nil, err
			}
		}
	case *AReleaseRq:
		pduType = TypeAReleaseRq
		err = enc.Write(encoding.Skip(4))
	case *AReleaseRp:
		pduType = TypeAReleaseRp
		err = enc.Write(encoding.Skip(4))
	case *AAbort:
		pduType = TypeAAbort
		err = enc.Write(encoding.Skip(2), &n.Source, &n.Reason)
	default:
		panic(fmt.Sprintf("Unknown PDU %v", pdu))
	}
	if err != nil {
		return nil, err
	}

	var header [6]byte // First 6 bytes of buf.
	header[0] = uint8(pduType)
	header[1] = 0 // Reserved.
	binary.BigEndian.PutUint32(header[2:6], uint32(enc.Len()))
	return append(header[:], enc.Bytes()...), nil
}

func writeSubItem(w *encoding.Writer, item any) error {
	switch val := item.(type) {
	case UserInformationMaximumLengthItem:
		return w.Write(ItemTypeUserInformationMaximumLength, encoding.Skip(1), uint16(4), val.MaximumLengthReceived)
	case AsynchronousOperationsWindowSubItem:
		return w.Write(ItemTypeAsynchronousOperationsWindow, encoding.Skip(1), uint16(4), val.MaxOpsPerformed, val.MaxOpsInvoked)
	case RoleSelectionSubItem:
		return w.Write(ItemTypeRoleSelection, encoding.Skip(1), uint16(2+len(val.SOPClassUID)+1*2), uint16(len(val.SOPClassUID)), val.SCURole, val.SCPRole)
	case SubItemUnsupported:
		return encSubItemWithName(w, ItemType(val.Type), string(val.Data))
	case ImplementationClassUIDSubItem:
		return encSubItemWithName(w, ItemTypeImplementationClassUID, val.Name)
	case ImplementationVersionNameSubItem:
		return encSubItemWithName(w, ItemTypeImplementationVersionName, val.Name)
	case ApplicationContextItem:
		return encSubItemWithName(w, ItemTypeApplicationContext, val.Name)
	case AbstractSyntaxSubItem:
		return encSubItemWithName(w, ItemTypeAbstractSyntax, val.Name)
	case TransferSyntaxSubItem:
		return encSubItemWithName(w, ItemTypeTransferSyntax, val.Name)
	case UserInformationItem:
		enc := encoding.NewWriter(binary.BigEndian)
		for _, s := range val.Items {
			if err := writeSubItem(enc, s); err != nil {
				return err
			}
		}
		return w.Write(ItemTypeUserInformation, encoding.Skip(1), uint16(enc.Len()), enc.Bytes())
	case PresentationContextItem:
		enc := encoding.NewWriter(binary.BigEndian)
		for _, s := range val.Items {
			if err := writeSubItem(enc, s); err != nil {
				return err
			}
		}
		return w.Write(val.Type, encoding.Skip(1), uint16(4+enc.Len()), val.ContextID, encoding.Skip(3), enc.Bytes())
	case PresentationDataValueItem:
		var header byte
		if val.Command {
			header |= 0b01
		}
		if val.Last {
			header |= 0b10
		}
		return w.Write(uint32(2+len(val.Value)), val.ContextID, header, val.Value)
	default:
		return fmt.Errorf("cannot write sub item type %T", item)
	}
}

func encSubItemWithName(w *encoding.Writer, itemType ItemType, name string) error {
	return w.Write(itemType, encoding.Skip(1), uint16(len(name)), []byte(name))
}

// padString pads the string with " " up to the given length.
func padString(v string, length int) string {
	if len(v) > length {
		return v[:length]
	}
	return v + strings.Repeat(" ", length-len(v))
}
