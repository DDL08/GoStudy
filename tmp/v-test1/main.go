package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func IpToUint32(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}
func Uint32ToIp(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
	// 创建一个长度为4的字节切片，表示IPv4地址的4个字节
	// 把nn这个32位无符号整数以大端字节序写入ip切片中
}
func main() {

	ip := net.ParseIP("192.168.1.1")
	fmt.Printf("IP: %d\n", ip)
	ipUint := IpToUint32(ip)
	fmt.Printf("IP: %d\n", ipUint)
	ip2 := Uint32ToIp(3232235777)
	fmt.Printf("IP: %d\n", ip2)
	// 结果是 net.IP{192, 168, 1, 1}
	//前 12 个字节 [0 0 0 0 0 0 0 0 0 0 0 0] 表示这是一个 IPv4 地址。
	//接下来的 4 个字节 [255 255 192 168 1 1] 表示 IPv4 地址 192.168.1.1。
	/*
		var u32 uint32 = 4294967295           // uint32 的最大值
		var u16 uint16 = 65535                // uint16 的最大值
		var u64 uint64 = 18446744073709551615 // uint64 的最大值
		fmt.Printf("uint32 最大值: %d\n", u32)
		fmt.Printf("uint16 最大值: %d\n", u16)
		fmt.Printf("uint64 最大值: %d\n", u64)*/

}
