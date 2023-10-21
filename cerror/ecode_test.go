package cerror

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsEcode(t *testing.T) {
	err := &Ecode{code: 12345}
	wrap1 := Wrap(err)

	assert.Equal(t, false, IsEcode(nil))
	assert.Equal(t, true, IsEcode(err))
	assert.Equal(t, true, IsEcode(wrap1))
	assert.Equal(t, false, IsEcode(errors.New("test")))
}

func TestGetEcode(t *testing.T) {
	err := &Ecode{code: 12345}
	wrap1 := Wrap(err)

	err1 := GetEcode(err)
	err2 := GetEcode(wrap1)
	err3 := GetEcode(nil)
	err4 := GetEcode(errors.New("test"))

	assert.Equal(t, true, err3 == nil)
	assert.Equal(t, 12345, err1.code)
	assert.Equal(t, 12345, err2.code)
	assert.Equal(t, true, err4 == nil)
}
