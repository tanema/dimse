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

func CreatePdata(contextID uint8, val []byte) []PDU {
	// DefaultMaxPDUSize
	return []PDU{&PDataTf{
		Items: []PresentationDataValueItem{{
			ContextID: contextID, // need to look up associated presentation context
			Command:   true,
			Last:      true,
			Value:     val,
		}},
	}}
}
