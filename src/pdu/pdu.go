package pdu

type (
	AAssociate struct {
		Type            Type
		ProtocolVersion uint16
		CalledAETitle   string
		CallingAETitle  string
		Items           []any
	}
	UserInformationItem                 struct{ Items []any }
	UserInformationMaximumLengthItem    struct{ MaximumLengthReceived uint32 }
	AsynchronousOperationsWindowSubItem struct {
		MaxOpsInvoked   uint16
		MaxOpsPerformed uint16
	}
	RoleSelectionSubItem struct {
		SOPClassUID string
		SCURole     uint8
		SCPRole     uint8
	}
	// Container for subitems that this package doesnt' support
	SubItemUnsupported struct {
		Type Type
		Data []byte
	}
	ImplementationClassUIDSubItem    struct{ Name string }
	ImplementationVersionNameSubItem struct{ Name string }
	ApplicationContextItem           struct{ Name string }
	AbstractSyntaxSubItem            struct{ Name string }
	TransferSyntaxSubItem            struct{ Name string }
	PresentationContextItem          struct {
		Type      ItemType
		ContextID uint8
		Result    PresentationContextResult
		Items     []any
	}
	AReleaseRq   struct{}
	AReleaseRp   struct{}
	AAssociateRj struct {
		Result RejectResultType
		Source SourceType
		Reason RejectReasonType
	}
	AAbort struct {
		Source SourceType
		Reason AbortReasonType
	}
	PDataTf struct {
		Items []PresentationDataValueItem
	}
	PresentationDataValueItem struct {
		ContextID uint8
		Command   bool // Bit 7 (LSB): 1 means command 0 means data
		Last      bool // Bit 6: 1 means last fragment. 0 means not last fragment.
		Value     []byte
	}
)
