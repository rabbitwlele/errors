package code

import (
	"fmt"
)

var (
	_codes = map[int]struct{}{} // register codes.
)

func New(e int, message string) Code {
	if _, ok := _codes[e]; ok {
		panic(fmt.Sprintf("code: %d already exist", e))
	}
	_codes[e] = struct{}{}
	return code{code: e, message: message}
}

type Code interface {
	Code() int
	Message() string
	Error() string
}

type code struct {
	code    int
	message string
}

func (e code) Code() int { return int(e.code) }

func (e code) Message() string { return e.message }

func (e code) Error() string { return e.message }
