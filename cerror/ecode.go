package cerror

import (
	"errors"
	"github.com/puresnr/go-cell/cast"
)

type ECode struct {
	Ecode int
}

func (e *ECode) Error() string { return cast.ToString(e.Ecode) }
func (e *ECode) Is(err error) bool {
	_, ok := err.(*ECode)
	return ok
}

var flagECode = new(ECode)

func IsECode(err error) bool {
	return errors.Is(err, flagECode)
}

func GetECode(err error) *ECode {
	e := new(ECode)
	if errors.As(err, &e) == true {
		return e
	}
	return nil
}
