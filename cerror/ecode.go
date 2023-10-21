package cerror

import (
	"errors"
	"github.com/puresnr/go-cell/cast"
)

type Ecode struct {
	code int
}

func New(code int) *Ecode { return &Ecode{code: code} }

func (e *Ecode) Error() string { return cast.ToString(e.code) }
func (e *Ecode) Is(err error) bool {
	_, ok := err.(*Ecode)
	return ok
}

func (e *Ecode) Code() int { return e.code }

var flagEcode = New(0)

func IsEcode(err error) bool {
	return errors.Is(err, flagEcode)
}

func GetEcode(err error) *Ecode {
	e := new(Ecode)
	if errors.As(err, &e) == true {
		return e
	}
	return nil
}
