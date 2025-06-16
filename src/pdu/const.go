package pdu

//go:generate stringer -type Type -trimprefix Type
//go:generate stringer -type ItemType -trimprefix ItemType
//go:generate stringer -type PresentationContextResult -trimprefix PresentationContext
//go:generate stringer -type RejectResultType -trimprefix ResultRejected
//go:generate stringer -type RejectReasonType -trimprefix RejectReason
//go:generate stringer -type SourceType -trimprefix SourceUL
//go:generate stringer -type AbortReasonType -trimprefix AbortReason
type (
	Type                      uint8
	ItemType                  uint8
	PresentationContextResult uint8
	RejectResultType          uint8
	RejectReasonType          uint8
	SourceType                uint8
	AbortReasonType           uint8
)

const (
	TypeAAssociateRq Type = 1 // A_ASSOCIATE_RQ association request
	TypeAAssociateAc Type = 2 // A_ASSOCIATE_AC association accepted
	TypeAAssociateRj Type = 3 // A_ASSOCIATE_RJ association rejuected
	TypePDataTf      Type = 4 // P_DATA_TF      used once an association has been established to send DIMSE message data.
	TypeAReleaseRq   Type = 5 // A_RELEASE_RQ   disassociation request
	TypeAReleaseRp   Type = 6 // A_RELEASE_RP   disassociation response
	TypeAAbort       Type = 7 // A_ABORT        disassociation without response.

	// Possible Type field values for SubItem.
	ItemTypeApplicationContext           ItemType = 0x10
	ItemTypePresentationContextRequest   ItemType = 0x20
	ItemTypePresentationContextResponse  ItemType = 0x21
	ItemTypeAbstractSyntax               ItemType = 0x30
	ItemTypeTransferSyntax               ItemType = 0x40
	ItemTypeUserInformation              ItemType = 0x50
	ItemTypeUserInformationMaximumLength ItemType = 0x51
	ItemTypeImplementationClassUID       ItemType = 0x52
	ItemTypeAsynchronousOperationsWindow ItemType = 0x53
	ItemTypeRoleSelection                ItemType = 0x54
	ItemTypeImplementationVersionName    ItemType = 0x55

	PresentationContextAccepted                                    PresentationContextResult = 0
	PresentationContextUserRejection                               PresentationContextResult = 1
	PresentationContextProviderRejectionNoReason                   PresentationContextResult = 2
	PresentationContextProviderRejectionAbstractSyntaxNotSupported PresentationContextResult = 3
	PresentationContextProviderRejectionTransferSyntaxNotSupported PresentationContextResult = 4

	ResultRejectedPermanent RejectResultType = 1
	ResultRejectedTransient RejectResultType = 2

	RejectReasonNone                               RejectReasonType = 1
	RejectReasonApplicationContextNameNotSupported RejectReasonType = 2
	RejectReasonCallingAETitleNotRecognized        RejectReasonType = 3
	RejectReasonCalledAETitleNotRecognized         RejectReasonType = 7

	SourceULServiceUser                 SourceType = 1
	SourceULServiceProviderACSE         SourceType = 2
	SourceULServiceProviderPresentation SourceType = 3

	AbortReasonNotSpecified             AbortReasonType = 0
	AbortReasonUnexpectedPDU            AbortReasonType = 2
	AbortReasonUnrecognizedPDUParameter AbortReasonType = 3
	AbortReasonUnexpectedPDUParameter   AbortReasonType = 4
	AbortReasonInvalidPDUParameterValue AbortReasonType = 5
)
