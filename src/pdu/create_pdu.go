package pdu

import (
	"github.com/tanema/dimse/src/serviceobjectpair"
	"github.com/tanema/dimse/src/transfersyntax"
)

func CreateAssoc(aetitle string, chunkSize uint32, sopsClasses []serviceobjectpair.UID, transfersyntaxes []transfersyntax.UID) (PDU, *ContextManager) {
	assoc := &AAssociate{
		Type:            TypeAAssociateRq,
		ProtocolVersion: CurrentProtocolVersion,
		CalledAETitle:   "anon-called-ae",
		CallingAETitle:  aetitle,
		Items:           []SubItem{&ApplicationContextItem{Name: DICOMApplicationContextItemName}},
	}

	cm := NewContextManager()
	for _, sop := range sopsClasses {
		assoc.Items = append(assoc.Items, cm.Add(sop, transfersyntaxes).ToPCI())
	}

	assoc.Items = append(assoc.Items,
		&UserInformationItem{
			Items: []SubItem{
				&UserInformationMaximumLengthItem{chunkSize},
				&ImplementationClassUIDSubItem{ImplementationClassUID},
				&ImplementationVersionNameSubItem{ImplementationName},
			},
		})
	return assoc, cm
}

func CreateRelease() PDU {
	return &AReleaseRq{}
}

func CreateAbort() PDU {
	return &AAbort{
		Source: SourceULServiceUser,
		Reason: AbortReasonNotSpecified,
	}
}

func CreatePdata(ctxID uint8, cmd bool, data []byte) []PDU {
	var pdus []PDU
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
			Items: []PresentationDataValueItem{
				{
					ContextID: ctxID,
					Command:   cmd,
					Last:      lastChunk, // Set later.
					Value:     chunk,
				},
			},
		})
	}
	return pdus
}
