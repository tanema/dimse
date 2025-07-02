package pdu

import (
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/tanema/dimse/src/defn/item"
	"github.com/tanema/dimse/src/encoding"
)

// EncodePDU serializes "pdu" into []byte.
func EncodePDU(pdu any, bo binary.ByteOrder) ([]byte, error) {
	enc := encoding.NewWriter(bo)
	var pduType Type
	var err error
	switch n := pdu.(type) {
	case *AAssociate:
		pduType = n.Type
		err = encodeAssociate(enc, bo, n)
	case *AAssociateRj:
		pduType = TypeAAssociateRj
		err = enc.WriteParts(encoding.Skip(1), n.Result, n.Source, n.Reason)
	case *PDataTf:
		pduType = TypePDataTf
		err = encodePData(enc, n)
	case *AReleaseRq:
		pduType = TypeAReleaseRq
		err = enc.WriteParts(encoding.Skip(4))
	case *AReleaseRp:
		pduType = TypeAReleaseRp
		err = enc.WriteParts(encoding.Skip(4))
	case *AAbort:
		pduType = TypeAAbort
		err = enc.WriteParts(encoding.Skip(2), &n.Source, &n.Reason)
	default:
		return nil, fmt.Errorf("Unknown PDU %v", pdu)
	}
	if err != nil {
		return nil, err
	}
	final := encoding.NewWriter(bo)
	final.WriteParts(uint8(pduType), uint8(0), uint32(enc.Len()))
	return append(final.Bytes(), enc.Bytes()...), nil
}

func encodePData(w *encoding.Writer, n *PDataTf) error {
	var header byte
	if n.Command {
		header |= 0b01
	}
	if n.Last {
		header |= 0b10
	}
	return w.WriteParts(uint32(2+len(n.Value)), n.ContextID, header, n.Value)
}

func encodeAssociate(w *encoding.Writer, bo binary.ByteOrder, n *AAssociate) error {
	if n.Type == 0 || n.CalledAETitle == "" || n.CallingAETitle == "" {
		return fmt.Errorf("Malformed associate")
	}
	if err := w.WriteParts(
		n.ProtocolVersion,
		encoding.Skip(2),
		[]byte(padString(n.CalledAETitle, 16)),
		[]byte(padString(n.CallingAETitle, 16)),
		encoding.Skip(32),
	); err != nil {
		return err
	}

	if err := w.WriteParts(item.ApplicationContext, encoding.Skip(1), uint16(len(n.ApplicationContext)), []byte(n.ApplicationContext)); err != nil {
		return err
	}

	for _, pitem := range n.PresentationItems {
		enc := encoding.NewWriter(bo)
		if err := enc.WriteParts(item.AbstractSyntax, encoding.Skip(1), uint16(len(pitem.AbstractSyntax)), []byte(pitem.AbstractSyntax)); err != nil {
			return err
		}
		for _, s := range pitem.TransferSyntaxes {
			if err := enc.WriteParts(item.TransferSyntax, encoding.Skip(1), uint16(len(s)), []byte(s)); err != nil {
				return err
			}
		}
		if err := w.WriteParts(item.PresentationContextRequest, encoding.Skip(1), uint16(4+enc.Len()), pitem.ContextID, encoding.Skip(3), enc.Bytes()); err != nil {
			return err
		}
	}

	enc := encoding.NewWriter(bo)
	if err := enc.WriteParts(item.UserInformationMaximumLength, encoding.Skip(1), uint16(4), n.MaximumLengthReceived); err != nil {
		return err
	} else if err := enc.WriteParts(item.ImplementationClassUID, encoding.Skip(1), uint16(len(n.ImplementationClassUID)), []byte(n.ImplementationClassUID)); err != nil {
		return err
	} else if err := enc.WriteParts(item.ImplementationVersionName, encoding.Skip(1), uint16(len(n.ImplementationVersionName)), []byte(n.ImplementationVersionName)); err != nil {
		return err
	} else if err := w.WriteParts(item.UserInformation, encoding.Skip(1), uint16(enc.Len()), enc.Bytes()); err != nil {
		return err
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
