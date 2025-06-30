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
		Command   bool // Bit 7 (LSB): 1 means command 0 means data
		Last      bool // Bit 6: 1 means last fragment. 0 means not last fragment.
		Value     []byte
	}
)
