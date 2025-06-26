package source

//go:generate stringer -type Type
type Type uint8

const (
	NotSpecified Type = iota
	ServiceUser
	ServiceProviderACSE
	ServiceProviderPresentation
)
