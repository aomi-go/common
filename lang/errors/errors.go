package errors

func NewIllegalArgumentError(msg string) *IllegalArgumentError {
	return &IllegalArgumentError{Msg: msg}
}

type IllegalArgumentError struct {
	Msg string
}

func (i *IllegalArgumentError) Error() string {
	return i.Msg
}
