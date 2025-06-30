package pdu

import (
	"github.com/tanema/dimse/src/defn/abort"
	"github.com/tanema/dimse/src/defn/serviceobjectpair"
	"github.com/tanema/dimse/src/defn/source"
	"github.com/tanema/dimse/src/defn/transfersyntax"
)

func CreateAssoc(localTitle, remoteTitle string, chunkSize uint32, sopsClasses []serviceobjectpair.UID, transfersyntaxes []transfersyntax.UID) (*AAssociate, *ContextManager) {
	cm := NewContextManager()
	pci := []PresentationContextItem{}
	for _, sop := range sopsClasses {
		pci = append(pci, cm.Add(sop, transfersyntaxes).ToPCI())
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

func CreateRelease() *AReleaseRq {
	return &AReleaseRq{}
}

func CreateAbort() *AAbort {
	return &AAbort{
		Source: source.ServiceUser,
		Reason: abort.NotSpecified,
	}
}

func CreatePdata(ctxID uint8, cmd bool, data []byte) []*PDataTf {
	var pdus []*PDataTf
	// two byte header overhead.
	maxChunkSize := int(DefaultMaxPDUSize - 8)
	for len(data) > 0 {
		chunkSize := len(data)
		if chunkSize > maxChunkSize {
			chunkSize = maxChunkSize
		}
		chunk := data[0:chunkSize]
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
