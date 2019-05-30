package test_go

import (
	"errors"
)


var(
	StatusDivideZero = errors.New("Divide Zero")
)

func Div(a float64, b float64) (float64, error)  {
	if b == 0 {
		return 0, StatusDivideZero
	}
	return a / b , nil
}