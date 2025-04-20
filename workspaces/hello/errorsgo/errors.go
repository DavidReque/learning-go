package errorsgo

import (
	//"errors"
	"fmt"
)

type DivisionError struct {
	Code int
	Msg string
}

func (d DivisionError) Error() string {
	return fmt.Sprintf("code %d: %s", d.Code, d.Msg)
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, DivisionError{Code: 2000, Msg: "No se puede dividir entre 0"}
	}

	return a/b, nil
}