package reject

//go:generate stringer -type Result
//go:generate stringer -type Reason
type (
	Result uint8
	Reason uint8
)

const (
	Permanent                          Result = 1
	Transient                          Result = 2
	None                               Reason = 1
	ApplicationContextNameNotSupported Reason = 2
	CallingAETitleNotRecognized        Reason = 3
	CalledAETitleNotRecognized         Reason = 7
)
