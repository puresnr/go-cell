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

func TestGetIP4LAN(t *testing.T) {
	ip, _ := GetIP4LAN()
	assert.Equal(t, "192.168.28.139", ip)
}

func TestGetIP4UintLAN(t *testing.T) {
	ip, _ := GetIP4LAN()
	intip, _ := GetIP4UintLAN()

	assert.Equal(t, IP4ToUint(ip), intip)
}
