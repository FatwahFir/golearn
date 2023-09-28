package exception

type NotFounError struct {
	Error string
}

func NewNotFoundError(err string) NotFounError {
	return NotFounError{
		Error: err,
	}
}
