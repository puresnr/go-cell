package cast

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"math"
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

func TestCastSafe(t *testing.T) {
	var (
		ui   uint
		ui32 uint32
		i32  int32
		err  error
		i    int
	)

	var mui64 uint64 = math.MaxUint64
	var mui32 uint32 = math.MaxUint32
	var mi32 = math.MaxInt32
	var mi64 = math.MaxInt64
	var minus = -1

	err = CastSafe(minus, &ui)
	assert.Equal(t, true, errors.Is(err, EOverflow))
	err = CastSafe(mi64, &ui32)
	assert.Equal(t, true, errors.Is(err, EOverflow))
	err = CastSafe(mi64, &ui)
	assert.Equal(t, mi64, int(ui))
	assert.Equal(t, nil, err)
	err = CastSafe(mi64, &i32)
	assert.Equal(t, true, errors.Is(err, EOverflow))
	err = CastSafe(mi32, &i32)
	assert.Equal(t, mi32, int(i32))
	assert.Equal(t, nil, err)
	err = CastSafe(mui64, &ui32)
	assert.Equal(t, true, errors.Is(err, EOverflow))
	err = CastSafe(mui32, &ui)
	assert.Equal(t, mui32, uint32(ui))
	assert.Equal(t, nil, err)
	err = CastSafe(mui64, &i32)
	assert.Equal(t, true, errors.Is(err, EOverflow))
	err = CastSafe(mui32, &i)
	assert.Equal(t, mui32, uint32(i))
	assert.Equal(t, nil, err)
}
