package commands

//go:generate stringer -type Kind
//go:generate stringer -type DataSetType
//go:generate stringer -type Status
type (
	Kind        int
	DataSetType int
	Status      int
)

const (
	CSTORERQ        Kind = 0x0001
	CSTORERSP       Kind = 0x8001
	CGETRQ          Kind = 0x0010
	CGETRSP         Kind = 0x8010
	CFINDRQ         Kind = 0x0020
	CFINDRSP        Kind = 0x8020
	CMOVERQ         Kind = 0x0021
	CMOVERSP        Kind = 0x8021
	CECHORQ         Kind = 0x0030
	CECHORSP        Kind = 0x8030
	NEVENTREPORTRQ  Kind = 0x0100
	NEVENTREPORTRSP Kind = 0x8100
	NGETRQ          Kind = 0x0110
	NGETRSP         Kind = 0x8110
	NSETRQ          Kind = 0x0120
	NSETRSP         Kind = 0x8120
	NACTIONRQ       Kind = 0x0130
	NACTIONRSP      Kind = 0x8130
	NCREATERQ       Kind = 0x0140
	NCREATERSP      Kind = 0x8140
	NDELETERQ       Kind = 0x0150
	NDELETERSP      Kind = 0x8150
	CCANCELRQ       Kind = 0x0FFF

	Null    DataSetType = 0x101
	NonNull DataSetType = 1

	Success               Status = 0
	Cancel                Status = 0xFE00
	SOPClassNotSupported  Status = 0x0112
	InvalidArgumentValue  Status = 0x0115
	InvalidAttributeValue Status = 0x0106
	InvalidObjectInstance Status = 0x0117
	UnrecognizedOperation Status = 0x0211
	NotAuthorized         Status = 0x0124
	Pending               Status = 0xff00
)
