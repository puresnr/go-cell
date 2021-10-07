package net

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIP4ToUint(t *testing.T) {
	assert.Equal(t, uint32(0), IP4ToUint("0.0"))
	assert.Equal(t, uint32(0), IP4ToUint("192.1.123.1.0"))
	assert.Equal(t, uint32(0), IP4ToUint("192.168.1.a"))
	assert.Equal(t, uint32(0), IP4ToUint("192.168.a.1"))
	assert.Equal(t, uint32(0), IP4ToUint("192.a.1.1"))
	assert.Equal(t, uint32(0), IP4ToUint("a.168.1.1"))
	assert.Equal(t, uint32(3232235777), IP4ToUint("192.168.1.1"))
}

func TestGetIP4Uint(t *testing.T) {
	ip, _ := GetIP4()
	intip, _ := GetIP4Uint()

	assert.Equal(t, IP4ToUint(ip), intip)
}
