package controllers

// Error this struct is for error.
type Error struct {
	Message string
}

// NewError this func is initializing error.
func NewError(err error) *Error {
	return &Error{
		Message: err.Error(),
	}
}
