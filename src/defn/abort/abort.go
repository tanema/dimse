package abort

//go:generate stringer -type Reason
type Reason uint8

const (
	NotSpecified Reason = iota
	Reserved
	UnexpectedPDU
	UnrecognizedPDUParameter
	UnexpectedPDUParameter
	InvalidPDUParameterValue
)
