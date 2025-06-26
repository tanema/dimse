package item

//go:generate stringer -type Type
type Type uint8

const (
	ApplicationContext           Type = 0x10
	PresentationContextRequest   Type = 0x20
	PresentationContextResponse  Type = 0x21
	AbstractSyntax               Type = 0x30
	TransferSyntax               Type = 0x40
	UserInformation              Type = 0x50
	UserInformationMaximumLength Type = 0x51
	ImplementationClassUID       Type = 0x52
	AsynchronousOperationsWindow Type = 0x53
	RoleSelection                Type = 0x54
	ImplementationVersionName    Type = 0x55
)
