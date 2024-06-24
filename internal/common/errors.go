package common

type NotFoundError struct {
	err error
}

func NewNotFoundError(err error) *NotFoundError {
	return &NotFoundError{err}
}

func (e *NotFoundError) Error() string {
	return e.err.Error()
}

type ConflictError struct {
	err error
}

func NewConflictError(err error) *ConflictError {
	return &ConflictError{err}
}

func (e *ConflictError) Error() string {
	return e.err.Error()
}
