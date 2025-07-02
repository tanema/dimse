package pdu

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/tanema/dimse/src/defn/item"
	"github.com/tanema/dimse/src/defn/serviceobjectpair"
	"github.com/tanema/dimse/src/defn/transfersyntax"
	"github.com/tanema/dimse/src/encoding"
)

type Reader struct {
	reader *encoding.Reader
}

func NewReader(in io.Reader) *Reader {
	return &Reader{reader: encoding.NewReader(in, binary.BigEndian, true)}
}

func (r *Reader) Next() (any, error) {
	var pduType Type
	var length uint32
	if err := r.reader.Read(&pduType, encoding.Skip(1), &length); err != nil {
		return nil, fmt.Errorf("ReadPDUs error reading pdu header: %v", err)
	}
	return r.pdu(pduType, int(length))
}

func (r *Reader) pdu(pduType Type, length int) (any, error) {
	r.reader.PushLimit(int(length))
	switch pduType {
	case TypeAAssociateRq, TypeAAssociateAc:
		return r.decodeAssociate()
	case TypePDataTf:
		return r.decodePData()
	case TypeAAssociateRj:
		assoc := &AAssociateRj{}
		return assoc, r.reader.Read(encoding.Skip(1), &assoc.Result, &assoc.Source, &assoc.Reason)
	case TypeAAbort:
		abort := &AAbort{}
		return abort, r.reader.Read(encoding.Skip(2), &abort.Source, &abort.Reason)
	case TypeAReleaseRq:
		return &AReleaseRq{}, r.reader.Read(encoding.Skip(4))
	case TypeAReleaseRp:
		return &AReleaseRp{}, r.reader.Read(encoding.Skip(4))
	default:
		return nil, fmt.Errorf("ReadPDU: unknown message type %d", pduType)
	}
}

func (r *Reader) decodeAssociate() (*AAssociate, error) {
	assoc := &AAssociate{
		Type: TypeAAssociateAc,
	}
	if err := r.reader.Read(&assoc.ProtocolVersion, encoding.Skip(2)); err != nil {
		return assoc, fmt.Errorf("error reading protocol version %v", err)
	} else if assoc.CalledAETitle, err = r.reader.String(16); err != nil {
		return assoc, fmt.Errorf("error reading called aetitle %v", err)
	} else if assoc.CallingAETitle, err = r.reader.String(16); err != nil {
		return assoc, fmt.Errorf("error reading calling aetitle %v", err)
	} else if err := r.reader.Read(encoding.Skip(8 * 4)); err != nil {
		return assoc, fmt.Errorf("error reading reserved data chunk %v", err)
	}

	var itemType item.Type
	var length uint16
	if err := r.reader.Read(&itemType, encoding.Skip(1), &length); err != nil {
		return nil, err
	} else if assoc.ApplicationContext, err = r.reader.String(int(length)); err != nil {
		return nil, err
	}

	for {
		if err := r.reader.Read(&itemType, encoding.Skip(1), &length); err != nil {
			return nil, err
		} else if itemType == item.PresentationContextResponse {
			v := PresentationContextItem{}
			r.reader.PushLimit(int(length))
			if err := r.reader.Read(&v.ContextID, encoding.Skip(1), &v.Result, encoding.Skip(1)); err != nil {
				return nil, err
			} else if v.ContextID%2 != 1 {
				return nil, fmt.Errorf("PresentationContextItem ID must be odd, but found %x", v.ContextID)
			}
			for {
				if err := r.reader.Read(&itemType, encoding.Skip(1), &length); err != nil {
					if errors.Is(err, io.EOF) {
						r.reader.PopLimit()
						assoc.PresentationItems = append(assoc.PresentationItems, v)
						break
					}
					return nil, err
				} else if itemType == item.AbstractSyntax {
					name, err := r.reader.String(int(length))
					if err != nil {
						return nil, err
					}
					v.AbstractSyntax = serviceobjectpair.UID(name)
				} else if itemType == item.TransferSyntax {
					name, err := r.reader.String(int(length))
					if err != nil {
						return nil, err
					}
					v.TransferSyntaxes = append(v.TransferSyntaxes, transfersyntax.UID(name))
				} else {
					break
				}
			}
		} else {
			break
		}
	}
	if itemType != item.UserInformation {
		return nil, fmt.Errorf("unexpected item type while expecting user information")
	} else if err := r.reader.Read(&itemType, encoding.Skip(1), &length); err != nil {
		return nil, err
	} else if length != 4 {
		return nil, fmt.Errorf("UserInformationMaximumLengthItem must be 4 bytes, but found %dB", length)
	} else if err := r.reader.Read(&assoc.MaximumLengthReceived); err != nil {
		return nil, err
	} else if err := r.reader.Read(&itemType, encoding.Skip(1), &length); err != nil {
		return nil, err
	} else if itemType != item.ImplementationClassUID {
		return nil, fmt.Errorf("expected ImplementationClassUID")
	} else if assoc.ImplementationClassUID, err = r.reader.String(int(length)); err != nil {
		return nil, err
	} else if err := r.reader.Read(&itemType, encoding.Skip(1), &length); err != nil {
		return nil, err
	} else if itemType != item.ImplementationVersionName {
		return nil, fmt.Errorf("expected ImplementationVersionName")
	} else if assoc.ImplementationVersionName, err = r.reader.String(int(length)); err != nil {
		return nil, err
	}
	return assoc, nil
}

func (r *Reader) decodePData() (*PDataTf, error) {
	pdtf := &PDataTf{}
	var length uint32
	var header uint8
	if err := r.reader.Read(&length, &pdtf.ContextID, &header); err != nil {
		if errors.Is(err, io.EOF) {
			return pdtf, nil
		}
		return pdtf, fmt.Errorf("error reading pdata header: %v", err)
	}
	pdtf.Command = (header&1 != 0)
	pdtf.Last = (header&2 != 0)
	pdtf.Value = make([]byte, int(length-2))
	if err := r.reader.Read(&pdtf.Value); err != nil {
		if errors.Is(err, io.EOF) {
			return pdtf, nil
		}
		return pdtf, fmt.Errorf("error reading pdata value: %v", err)
	}
	return pdtf, nil
}
