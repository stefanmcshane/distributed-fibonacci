package business

import (
	"fmt"
	"math/big"
)

const (
	ErrorInvalidOrdinal  = "Must supply valid ordinal - 0 or above"
	ErrorOverflow        = "Fibonacci value overflowed given arch limit for int"
	ErrorStorageAddError = "Unable to store new ordinal in database"
)

type FibError struct {
	Ordinal big.Int
	Message string
}

func (f FibError) Error() string {
	return fmt.Sprintf("error: %s || ordinal: %s", f.Message, f.Ordinal.String())
}

func NewError(ordinal big.Int, message string) FibError {
	return FibError{
		Ordinal: ordinal,
		Message: message,
	}
}
