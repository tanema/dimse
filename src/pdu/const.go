package pdu

//go:generate stringer -type Type -trimprefix Type
type Type uint8

const (
	// The app context for DICOM. The first item in the A-ASSOCIATE-RQ
	DICOMApplicationContextItemName        = "1.2.840.10008.3.1.1.1"
	ImplementationClassUID                 = "1.2.826.0.1.3680043.9.7133"
	ImplementationName                     = "GODICOM_1_1"
	CurrentProtocolVersion          uint16 = 1
	DefaultMaxPDUSize               uint32 = 4 << 20

	TypeAAssociateRq Type = 1 // A_ASSOCIATE_RQ association request
	TypeAAssociateAc Type = 2 // A_ASSOCIATE_AC association accepted
	TypeAAssociateRj Type = 3 // A_ASSOCIATE_RJ association rejuected
	TypePDataTf      Type = 4 // P_DATA_TF      used once an association has been established to send DIMSE message data.
	TypeAReleaseRq   Type = 5 // A_RELEASE_RQ   disassociation request
	TypeAReleaseRp   Type = 6 // A_RELEASE_RP   disassociation response
	TypeAAbort       Type = 7 // A_ABORT        disassociation without response.
)
