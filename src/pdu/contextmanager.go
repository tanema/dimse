package pdu

import (
	"fmt"

	"github.com/tanema/dimse/src/defn/serviceobjectpair"
	"github.com/tanema/dimse/src/defn/transfersyntax"
)

type (
	ContextManager struct {
		contextID   uint8
		idToContext map[uint8]*PresentationContext
		sopToID     map[serviceobjectpair.UID]uint8
	}
	PresentationContext struct {
		ContextID              uint8
		ServiceObjectPair      serviceobjectpair.UID
		TransferSyntaxes       []transfersyntax.UID
		Accepted               bool
		AcceptedTransferSyntax transfersyntax.UID
	}
)

func NewContextManager() *ContextManager {
	return &ContextManager{
		contextID:   1,
		idToContext: map[uint8]*PresentationContext{},
		sopToID:     map[serviceobjectpair.UID]uint8{},
	}
}

func (cm *ContextManager) Add(sop serviceobjectpair.UID, ts []transfersyntax.UID) *PresentationContext {
	pc := &PresentationContext{
		ContextID:         cm.contextID,
		ServiceObjectPair: sop,
		TransferSyntaxes:  ts,
	}
	cm.idToContext[cm.contextID] = pc
	cm.sopToID[sop] = cm.contextID
	cm.contextID += 2 // must be odd.
	// TODO protect id overflow
	return pc
}

func (cm *ContextManager) Accept(ctxID uint8, ts transfersyntax.UID) error {
	pc, found := cm.idToContext[ctxID]
	if !found {
		return fmt.Errorf("ctxID not register in context manager")
	}
	pc.Accepted = true
	pc.AcceptedTransferSyntax = ts
	return nil
}

func (cm *ContextManager) GetAccepted(sops ...serviceobjectpair.UID) (uint8, transfersyntax.UID, error) {
	for _, classUID := range sops {
		if pctx, err := cm.GetWithSOP(classUID); err == nil && pctx.Accepted {
			return pctx.ContextID, pctx.AcceptedTransferSyntax, nil
		}
	}
	return 0, "", fmt.Errorf("Could not find an associated presentation context item for command which means the server rejected the AffectedSOPClassUID you requested.")
}

func (cm *ContextManager) GetWithSOP(sop serviceobjectpair.UID) (*PresentationContext, error) {
	ctxID, found := cm.sopToID[sop]
	if !found {
		return nil, fmt.Errorf("sop %v not register in context manager", sop)
	}
	pc := cm.idToContext[ctxID]
	if !pc.Accepted {
		return nil, fmt.Errorf("sop %v not accepted during association", pc.ServiceObjectPair)
	}
	return pc, nil
}

func (p *PresentationContext) ToPCI() PresentationContextItem {
	return PresentationContextItem{
		ContextID:        p.ContextID,
		AbstractSyntax:   p.ServiceObjectPair,
		TransferSyntaxes: p.TransferSyntaxes,
	}
}
