// package pdu establishes the lower protocol level data structures for general
// dimse communication.
package pdu

import (
	"github.com/tanema/dimse/src/defn/abort"
	"github.com/tanema/dimse/src/defn/presentationctx"
	"github.com/tanema/dimse/src/defn/reject"
	"github.com/tanema/dimse/src/defn/serviceobjectpair"
	"github.com/tanema/dimse/src/defn/source"
	"github.com/tanema/dimse/src/defn/transfersyntax"
)

type (
	AAssociate struct {
		Type                      Type
		ProtocolVersion           uint16
		CalledAETitle             string
		CallingAETitle            string
		ApplicationContext        string
		PresentationItems         []PresentationContextItem
		MaximumLengthReceived     uint32
		ImplementationClassUID    string
		ImplementationVersionName string
	}
	PresentationContextItem struct {
		ContextID        uint8
		Result           presentationctx.Result
		AbstractSyntax   serviceobjectpair.UID
		TransferSyntaxes []transfersyntax.UID
	}
	AReleaseRq   struct{}
	AReleaseRp   struct{}
	AAssociateRj struct {
		Result reject.Result
		Source source.Type
		Reason reject.Reason
	}
	AAbort struct {
		Source source.Type
		Reason abort.Reason
	}
	PDataTf struct {
		ContextID uint8
		Command   bool
		Last      bool
		Value     []byte
	}
)

func CreateAssoc(localTitle, remoteTitle string, chunkSize uint32, sopsClasses []serviceobjectpair.UID, transfersyntaxes []transfersyntax.UID) (*AAssociate, *ContextManager) {
	cm := NewContextManager()
	pci := []PresentationContextItem{}
	for _, sop := range sopsClasses {
		pci = append(pci, PresentationContextItem{
			ContextID:        cm.Add(sop, transfersyntaxes).ContextID,
			AbstractSyntax:   sop,
			TransferSyntaxes: transfersyntaxes,
		})
	}
	return &AAssociate{
		Type:                      TypeAAssociateRq,
		ProtocolVersion:           CurrentProtocolVersion,
		CallingAETitle:            localTitle,
		CalledAETitle:             remoteTitle,
		ApplicationContext:        DICOMApplicationContextItemName,
		PresentationItems:         pci,
		MaximumLengthReceived:     chunkSize,
		ImplementationClassUID:    ImplementationClassUID,
		ImplementationVersionName: ImplementationName,
	}, cm
}

func CreatePdata(ctxID uint8, cmd bool, maxChunkSize int, data []byte) []*PDataTf {
	var pdus []*PDataTf
	maxChunkSize = maxChunkSize - 8
	for len(data) > 0 {
		chunkSize := min(maxChunkSize, len(data))
		chunk := data[:chunkSize]
		data = data[chunkSize:]
		lastChunk := len(data) == 0
		pdus = append(pdus, &PDataTf{
			ContextID: ctxID,
			Command:   cmd,
			Last:      lastChunk,
			Value:     chunk,
		})
	}
	return pdus
}
