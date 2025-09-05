//go:build 386 || amd64
// +build 386 amd64

/*
针对 386 和 amd64 架构（即常见的x86 32位和64位桌面/服务器架构）编写的全局配置
*/

package global

// 除去PacketHeader之外的部分最大允许使用的内存
// 防止单个数据包占用内存过大
const MAX_PACKET_SIZE = 20480

// 一次最多可以承受7168个连接
const TCP_MAX_CONNECTION = 1024

// 命令&数据通道最大缓冲区大小
const BUFFER_SIZE = 1024



/* ----------------------- 👇 --------------------------- */
// 协议数据分隔符,为了在这种字节流中分清每条消息的起止位置，常会在每条消息中加入一个特定的“分隔符”字符串
var PROTOCOL_SEPARATOR = "DDDD"

// 协议特征, 用于在端口重用时鉴别
var PROTOCOL_FEATURE = "DDDDDDDD"

// 密钥
var SECRET_KEY []byte
