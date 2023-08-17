package cast

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToString(t *testing.T) {
	assert.Equal(t, "-50000", ToString(-50000))
}

func TestToString_u(t *testing.T) {
	var ti uint32 = 50000
	assert.Equal(t, "50000", ToString_u(ti))
}

func TestToStringBase(t *testing.T) {
	assert.Equal(t, "10", ToStringBase(8, 8))
}

func TestToStringBase_u(t *testing.T) {
	var ti uint32 = 8
	assert.Equal(t, "10", ToStringBase_u(ti, 8))
}
