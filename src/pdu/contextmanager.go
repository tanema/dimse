package pdu

import "fmt"

type (
	ContextManager struct {
		contextID   uint8
		idToContext map[uint8]*PresentationContext
		sopToID     map[string]uint8
	}
	PresentationContext struct {
		ContextID         uint8
		ServiceObjectPair string
		TransferSyntaxes  []string
		Accepted          bool
	}
)

func NewContextManager() *ContextManager {
	return &ContextManager{
		contextID:   1,
		idToContext: map[uint8]*PresentationContext{},
		sopToID:     map[string]uint8{},
	}
}

func (cm *ContextManager) Add(sop string, ts []string) *PresentationContext {
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

func (cm *ContextManager) Accept(ctxID uint8) error {
	pc, found := cm.idToContext[ctxID]
	if !found {
		return fmt.Errorf("ctxID not register in context manager")
	}
	pc.Accepted = true
	return nil
}

func (cm *ContextManager) GetWithSOP(sop string) (*PresentationContext, error) {
	ctxID, found := cm.sopToID[sop]
	if !found {
		return nil, fmt.Errorf("sop %v not register in context manager", sop)
	}
	return cm.GetWithCtxID(ctxID)
}

func (cm *ContextManager) GetWithCtxID(id uint8) (*PresentationContext, error) {
	pc := cm.idToContext[id]
	if !pc.Accepted {
		return nil, fmt.Errorf("sop %v not accepted during association", pc.ServiceObjectPair)
	}
	return pc, nil
}

func (p *PresentationContext) ToPCI() *PresentationContextItem {
	syntaxItems := []SubItem{&AbstractSyntaxSubItem{Name: p.ServiceObjectPair}}
	for _, syntaxUID := range p.TransferSyntaxes {
		syntaxItems = append(syntaxItems, &TransferSyntaxSubItem{Name: syntaxUID})
	}
	return &PresentationContextItem{
		Type:      ItemTypePresentationContextRequest,
		ContextID: p.ContextID,
		Items:     syntaxItems,
	}
}
