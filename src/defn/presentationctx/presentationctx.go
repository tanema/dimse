package presentationctx

//go:generate stringer -type Result
type Result uint8

const (
	Accepted                                    Result = 0
	UserRejection                               Result = 1
	ProviderRejectionNoReason                   Result = 2
	ProviderRejectionAbstractSyntaxNotSupported Result = 3
	ProviderRejectionTransferSyntaxNotSupported Result = 4
)
