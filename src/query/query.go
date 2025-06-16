package query

//go:generate stringer -type Level
type (
	Level int
)

const (
	Patient Level = iota
	Study
	Series
)
