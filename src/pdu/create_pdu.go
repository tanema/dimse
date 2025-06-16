package pdu

func CreateAssoc(sopsClasses []string, transfersyntaxes []string) (PDU, *ContextManager) {
	assoc := &AAssociate{
		Type:            TypeAAssociateRq,
		ProtocolVersion: CurrentProtocolVersion,
		CalledAETitle:   "anon-called-ae",
		CallingAETitle:  "anon-calling-ae",
		Items:           []SubItem{&ApplicationContextItem{Name: DICOMApplicationContextItemName}},
	}

	cm := NewContextManager()
	for _, sop := range sopsClasses {
		assoc.Items = append(assoc.Items, cm.Add(sop, transfersyntaxes).ToPCI())
	}

	assoc.Items = append(assoc.Items,
		&UserInformationItem{
			Items: []SubItem{
				&UserInformationMaximumLengthItem{DefaultMaxPDUSize},
				&ImplementationClassUIDSubItem{"1.2.826.0.1.3680043.9.7133"},
				&ImplementationVersionNameSubItem{"GODICOM_1_1"},
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
