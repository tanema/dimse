package pdu

import (
	"github.com/tanema/dimse/src/defn/abort"
	"github.com/tanema/dimse/src/defn/serviceobjectpair"
	"github.com/tanema/dimse/src/defn/source"
	"github.com/tanema/dimse/src/defn/transfersyntax"
)

func CreateAssoc(localTitle, remoteTitle string, chunkSize uint32, sopsClasses []serviceobjectpair.UID, transfersyntaxes []transfersyntax.UID) (*AAssociate, *ContextManager) {
	assoc := &AAssociate{
		Type:            TypeAAssociateRq,
		ProtocolVersion: CurrentProtocolVersion,
		CallingAETitle:  localTitle,
		CalledAETitle:   remoteTitle,
		Items:           []any{ApplicationContextItem{Name: DICOMApplicationContextItemName}},
	}

	cm := NewContextManager()
	for _, sop := range sopsClasses {
		assoc.Items = append(assoc.Items, cm.Add(sop, transfersyntaxes).ToPCI())
	}

	assoc.Items = append(assoc.Items,
		UserInformationItem{
			Items: []any{
				UserInformationMaximumLengthItem{chunkSize},
				ImplementationClassUIDSubItem{ImplementationClassUID},
				ImplementationVersionNameSubItem{ImplementationName},
			},
		})
	return assoc, cm
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
