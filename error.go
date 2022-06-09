package golisper

type GoLisperErrorType int

const (
	ErrUnknownToken GoLisperErrorType = iota
	ErrUnknownValue
)

type GoLisperError struct {
	Type   GoLisperErrorType
	Reason string
}

func newError(t GoLisperErrorType, reason string) *GoLisperError {
	return &GoLisperError{Type: t, Reason: reason}
}

func (e *GoLisperError) Error() string {
	return e.Reason
}
