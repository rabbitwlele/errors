package ecode

import (
	"fmt"
)

var (
	_codes = map[int]struct{}{} // register codes.
)

// New new a ecode.Code by int value.
// NOTE: ecode must unique in global, the New will check repeat and then panic.
func New(e int, message string) Ecode {
	if _, ok := _codes[e]; ok {
		panic(fmt.Sprintf("code: %d already exist", e))
	}
	_codes[e] = struct{}{}
	return code{code: e, message: message}
}

// Code ecode error interface which has a code & message.
type Ecode interface {
	// Code get error code.
	Code() int
	// Message get code message.
	Message() string
}

type code struct {
	code    int
	message string
}

func (e code) Code() int { return int(e.code) }

func (e code) Message() string { return e.message }
