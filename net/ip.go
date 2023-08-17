package net

import (
	"fmt"
	"github.com/puresnr/go-cell/cast"
	"github.com/puresnr/go-cell/cerror"
	"net"
	"strconv"
	"strings"
)

// 因为 映客 GO 版本较低，不支持这个函数，所以拷贝了一份出来
// isPrivateIP reports whether ip is a private address, according to
// RFC 1918 (IPv4 addresses) and RFC 4193 (IPv6 addresses).
func isPrivateIP(ip net.IP) bool {
	if ip4 := ip.To4(); ip4 != nil {
		// Following RFC 1918, Section 3. Private Address Space which says:
		//   The Internet Assigned Numbers Authority (IANA) has reserved the
		//   following three blocks of the IP address space for private internets:
		//     10.0.0.0        -   10.255.255.255  (10/8 prefix)
		//     172.16.0.0      -   172.31.255.255  (172.16/12 prefix)
		//     192.168.0.0     -   192.168.255.255 (192.168/16 prefix)
		return ip4[0] == 10 ||
			(ip4[0] == 172 && ip4[1]&0xf0 == 16) ||
			(ip4[0] == 192 && ip4[1] == 168)
	}
	// Following RFC 4193, Section 8. IANA Considerations which says:
	//   The IANA has assigned the FC00::/7 prefix to "Unique Local Unicast".
	return len(ip) == net.IPv6len && ip[0]&0xfe == 0xfc
}

// Deprecated : 使用 GetIP4LAN(), 因为某些情况下，该函数会返回 B 类地址
func GetLocalIP() (string, error) {
	// 获取所有网卡
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	// 取第一个非lo的网卡IP
	for _, addr := range addrs {
		// 这个网络地址是IP地址: ipv4, ipv6
		ipNet, ok := addr.(*net.IPNet)
		if !ok || ipNet.IP.IsLoopback() {
			continue
		}
		// 跳过IPV6
		if ipNet.IP.To4() != nil {
			return ipNet.IP.String(), nil
		}
	}
	return "", fmt.Errorf("ip not found")
}

// GetIP4LAN 获取本机的局域网 IP4 地址， 以 string 形式返回，获取不到时，返回 ""
func GetIP4LAN() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", cerror.Wrap(err)
	}

	for _, value := range addrs {
		if ipnet, ok := value.(*net.IPNet); ok && isPrivateIP(ipnet.IP) {
			if ip4 := ipnet.IP.To4(); ip4 != nil {
				return cast.ToString_u(uint32(ip4[0])) + "." + cast.ToString_u(uint32(ip4[1])) + "." +
					cast.ToString_u(uint32(ip4[2])) + "." + cast.ToString_u(uint32(ip4[3])), nil
			}
		}
	}

	return "", nil
}

// GetIP4LANInt 获取本机的局域网 IP4 地址，并将其转换成一个 uint32 整数返回, 获取不到时，返回 0
func GetIP4LANInt() (uint32, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return 0, cerror.Wrap(err)
	}

	for _, value := range addrs {
		if ipnet, ok := value.(*net.IPNet); ok && isPrivateIP(ipnet.IP) {
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

// GetAvailableTcpPort : 获取一个可用的tcp端口, 返回0表示没有可用端口
func GetAvailableTcpPort() (uint32, error) {
	address, err := net.ResolveTCPAddr("tcp", "0.0.0.0:0")
	if err != nil {
		return 0, cerror.Wrap(err)
	}

	listener, err := net.ListenTCP("tcp", address)
	if err != nil {
		return 0, cerror.Wrap(err)
	}
	if listener == nil {
		return 0, nil
	}

	defer listener.Close()

	if ta, ok := listener.Addr().(*net.TCPAddr); !ok {
		return 0, nil
	} else {
		return uint32(ta.Port), nil
	}
}
