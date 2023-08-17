package csync

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGoWait(t *testing.T) {
	var test1, test2 int
	GoWait(func() {
		time.Sleep(1 * time.Second)
		test1 = 100
	}, func() { test2 = 200 })
	assert.Equal(t, 100, test1)
	assert.Equal(t, 200, test2)
}
