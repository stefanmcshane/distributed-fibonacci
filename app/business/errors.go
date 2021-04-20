package business

import "fmt"

const (
	ErrorInvalidOrdinal  = "Must supply valid ordinal - 0 or above"
	ErrorOverflow        = "Fibonacci value overflowed given arch limit for int"
	ErrorStorageAddError = "Unable to store new ordinal in database"
)

type FibError struct {
	Ordinal int64
	Message string
}

func (f FibError) Error() string {
	return fmt.Sprintf("error: %s || ordinal: %d", f.Message, f.Ordinal)
}

func NewError(ordinal int64, message string) FibError {
	return FibError{
		Ordinal: ordinal,
		Message: message,
	}
}
