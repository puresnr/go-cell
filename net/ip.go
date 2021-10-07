// Package net 该文件下包含一些 IP 的相关操作
package net

import (
	"github.com/puresnr/go-cell/cast"
	"net"
	"strconv"
	"strings"
)

// GetIP4 : 获取本机的IP4地址， 以 string 形式返回， 获取不到时，返回 ""
func GetIP4() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, value := range addrs {
		if ipnet, ok := value.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ip4 := ipnet.IP.To4(); ip4 != nil {
				return cast.ToString_32u(uint32(ip4[0])) + "." + cast.ToString_32u(uint32(ip4[1])) + "." + cast.ToString_32u(uint32(ip4[2])) + "." + cast.ToString_32u(uint32(ip4[3])), nil
			}
		}
	}

	return "", nil
}

// GetIP4Uint : 获取本机的IP4地址，并将其转换成一个 uint32 整数返回, 获取不到时，返回 0
func GetIP4Uint() (uint32, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return 0, err
	}

	for _, value := range addrs {
		if ipnet, ok := value.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ip4 := ipnet.IP.To4(); ip4 != nil {
				return (uint32(ip4[0]) << 24) | (uint32(ip4[1]) << 16) | (uint32(ip4[2]) << 8) | uint32(ip4[3]), nil
			}
		}
	}

	return 0, nil
}

// IP4ToUint : 把一个指定的 ip4 地址转换为一个 uint32 整数，不是一个合法 ip4 地址时，返回 0
func IP4ToUint(ip string) uint32 {
	ints := strings.Split(ip, ".")
	if len(ints) != 4 {
		return 0
	}

	i, err := strconv.Atoi(ints[0])
	if err != nil {
		return 0
	}
	ret := uint32(i) << 24

	i, err = strconv.Atoi(ints[1])
	if err != nil {
		return 0
	}
	ret |= uint32(i) << 16

	i, err = strconv.Atoi(ints[2])
	if err != nil {
		return 0
	}
	ret |= uint32(i) << 8

	i, err = strconv.Atoi(ints[3])
	if err != nil {
		return 0
	}
	return ret | uint32(i)
}
