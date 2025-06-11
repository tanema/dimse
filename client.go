package dimse

import (
	"bytes"
	"fmt"
	"net"
	"sync/atomic"

	"github.com/tanema/dimse/sops"
	"github.com/tanema/dimse/transfersyntax"
)

const DefaultMaxPDUSize uint32 = 4 << 20

type Client struct {
	addr      string
	errors    chan error
	events    chan PDU
	conn      net.Conn
	msgID     int32
	connected bool
}

func Connect(addr string) (*Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	client := &Client{
		conn:      conn,
		connected: false,
		errors:    make(chan error),
		events:    make(chan PDU),
	}
	go client.listen()
	return client, client.associate(sops.StorageClasses, transfersyntax.Standard)
}

func (c *Client) nextMsgID() int32 {
	return atomic.AddInt32(&c.msgID, 1)
}

func (c *Client) listen() {
	for {
		data := make([]byte, 4096)
		n, err := c.conn.Read(data)
		if err != nil {
			continue
		}
		if n > 0 {
			pdu, err := ReadPDU(bytes.NewBuffer(data[:n]))
			if err != nil {
				c.errors <- err
			} else {
				c.handlePDU(pdu)
			}
		}
	}
}

func (c *Client) handlePDU(pdu PDU) {
	// handle association requests
	switch p := pdu.(type) {
	case *AAssociate:
		if p.Type == TypeAAssociateAc {
			c.connected = true
		}
	case *AAssociateRj:
		c.connected = false
		c.errors <- fmt.Errorf("Association rejected")
	case *PDataTf:
		c.errors <- fmt.Errorf("server requested dicom data from client")
	case *AReleaseRq:
		c.connected = false
		rel, _ := EncodePDU(&AReleaseRp{})
		go c.send(rel)
	case *AReleaseRp, *AAbort:
		c.connected = false
	default:
		c.events <- pdu
	}
}

func (c *Client) Errors() chan error { return c.errors }
func (c *Client) Events() chan PDU   { return c.events }

func (c *Client) Send(body []byte) error {
	if !c.connected {
		return fmt.Errorf("client not yet connected, association not complete.")
	}
	return c.send(body)
}

func (c *Client) send(body []byte) error {
	_, err := c.conn.Write(body)
	return err
}

func (c *Client) Close() error {
	close(c.events)
	close(c.errors)
	return c.conn.Close()
}

func (c *Client) associate(sopsClasses []string, transfersyntaxes []string) error {
	pdu := &AAssociate{
		Type:            TypeAAssociateRq,
		ProtocolVersion: CurrentProtocolVersion,
		CalledAETitle:   "anon-called-ae",
		CallingAETitle:  "anon-calling-ae",
		Items:           []SubItem{&ApplicationContextItem{Name: DICOMApplicationContextItemName}},
	}

	var contextID byte = 1
	for _, sop := range sopsClasses {
		syntaxItems := []SubItem{&AbstractSyntaxSubItem{Name: sop}}
		for _, syntaxUID := range transfersyntaxes {
			syntaxItems = append(syntaxItems, &TransferSyntaxSubItem{Name: syntaxUID})
		}
		pdu.Items = append(pdu.Items, &PresentationContextItem{
			Type:      ItemTypePresentationContextRequest,
			ContextID: contextID,
			Result:    0, // must be zero for request
			Items:     syntaxItems,
		})
		contextID += 2 // must be odd.
	}
	pdu.Items = append(pdu.Items,
		&UserInformationItem{
			Items: []SubItem{
				&UserInformationMaximumLengthItem{DefaultMaxPDUSize},
				&ImplementationClassUIDSubItem{"1.2.826.0.1.3680043.9.7133"},
				&ImplementationVersionNameSubItem{"GODICOM_1_1"},
			},
		})

	data, err := EncodePDU(pdu)
	if err != nil {
		return err
	}
	return c.send(data)
}

// func (c *Client) CEcho() error {
//	return c.send(&CEcho{
//		MessageID:          int(c.nextMsgID()),
//		CommandDataSetType: CommandDataSetTypeNull,
//	})
// }

// func (su *ServiceUser) CStore(ds *dicom.DataSet) error {
//	var sopClassUID string
//	if sopClassUIDElem, err := ds.FindElementByTag(dicomtag.MediaStorageSOPClassUID); err != nil {
//		return err
//	} else if sopClassUID, err = sopClassUIDElem.GetString(); err != nil {
//		return err
//	}
//	context, err := su.cm.lookupByAbstractSyntaxUID(sopClassUID)
//	if err != nil {
//		return err
//	}
//	cs, err := su.disp.newCommand(su.cm, context)
//	if err != nil {
//		return err
//	}
//	if err != nil {
//		return err
//	}
//	defer su.disp.deleteCommand(cs)
//	return runCStoreOnAssociation(cs.upcallCh, su.disp.downcallCh, su.cm, cs.messageID, ds)
// }

// func (su *ServiceUser) CFind(qrLevel QRLevel, filter []*dicom.Element) error {
//	context, payload, err := encodeQRPayload(qrOpCFind, qrLevel, filter, su.cm)
//	if err != nil {
//		return err
//	}
//	cs.sendMessage(
//		&dimse.CFind{
//			AffectedSOPClassUID: context.abstractSyntaxUID,
//			MessageID:           cs.messageID,
//			CommandDataSetType:  dimse.CommandDataSetTypeNonNull,
//		},
//		payload)
//	return nil
// }

// func (su *ServiceUser) CGet(qrLevel QRLevel, filter []*dicom.Element) error {
//	context, payload, err := encodeQRPayload(qrOpCGet, qrLevel, filter, su.cm)
//	if err != nil {
//		return err
//	}
//	cs.sendMessage(
//		&dimse.CGet{
//			AffectedSOPClassUID: context.abstractSyntaxUID,
//			MessageID:           cs.messageID,
//			CommandDataSetType:  dimse.CommandDataSetTypeNonNull,
//		},
//		payload)
// }

// func encodeQRPayload(opType qrOpType, qrLevel QRLevel, filter []*dicom.Element, cm *contextManager) (contextManagerEntry, []byte, error) {
//	var sopClassUID string
//	var qrLevelString string
//	switch qrLevel {
//	case QRLevelPatient:
//		switch opType {
//		case qrOpCFind:
//			sopClassUID = dicomuid.PatientRootQRFind
//		case qrOpCGet:
//			sopClassUID = dicomuid.PatientRootQRGet
//		case qrOpCMove:
//			sopClassUID = dicomuid.PatientRootQRMove
//		}
//		qrLevelString = "PATIENT"
//	case QRLevelStudy, QRLevelSeries:
//		switch opType {
//		case qrOpCFind:
//			sopClassUID = dicomuid.StudyRootQRFind
//		case qrOpCGet:
//			sopClassUID = dicomuid.StudyRootQRGet
//		case qrOpCMove:
//			sopClassUID = dicomuid.StudyRootQRMove
//		}
//		qrLevelString = "STUDY"
//		if qrLevel == QRLevelSeries {
//			qrLevelString = "SERIES"
//		}
//	default:
//		return contextManagerEntry{}, nil, fmt.Errorf("Invalid C-FIND QR lever: %d", qrLevel)
//	}

//	// Translate qrLevel to the sopclass and QRLevel elem.
//	// Encode the C-FIND DIMSE command.
//	context, err := cm.lookupByAbstractSyntaxUID(sopClassUID)
//	if err != nil {
//		// This happens when the user passed a wrong sopclass list in
//		// A-ASSOCIATE handshake.
//		return context, nil, err
//	}

//	// Encode the data payload containing the filtering conditions.
//	dataEncoder := dicomio.NewBytesEncoderWithTransferSyntax(context.transferSyntaxUID)
//	foundQRLevel := false
//	for _, elem := range filter {
//		if elem.Tag == dicomtag.QueryRetrieveLevel {
//			foundQRLevel = true
//		}
//		dicom.WriteElement(dataEncoder, elem)
//	}
//	if !foundQRLevel {
//		elem := dicom.MustNewElement(dicomtag.QueryRetrieveLevel, qrLevelString)
//		dicom.WriteElement(dataEncoder, elem)
//	}
//	if err := dataEncoder.Error(); err != nil {
//		return context, nil, err
//	}
//	return context, dataEncoder.Bytes(), err
// }
