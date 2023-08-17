package cerror

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestECode_Is(t *testing.T) {
	err := &ECode{Ecode: 12345}

	assert.Equal(t, true, err.Is(flagECode))
	assert.Equal(t, false, errors.Is(err, errors.New("test")))
	assert.Equal(t, false, errors.Is(err, nil))
}

func TestIsECode(t *testing.T) {
	err := &ECode{Ecode: 12345}
	wrap1 := Wrap(err)

	assert.Equal(t, false, IsECode(nil))
	assert.Equal(t, true, IsECode(err))
	assert.Equal(t, true, IsECode(wrap1))
	assert.Equal(t, false, IsECode(errors.New("test")))
}

func TestGetECode(t *testing.T) {
	err := &ECode{Ecode: 12345}
	wrap1 := Wrap(err)

	err1 := GetECode(err)
	err2 := GetECode(wrap1)
	err3 := GetECode(nil)
	err4 := GetECode(errors.New("test"))

	assert.Equal(t, true, err3 == nil)
	assert.Equal(t, 12345, err1.Ecode)
	assert.Equal(t, 12345, err2.Ecode)
	assert.Equal(t, true, err4 == nil)
}
