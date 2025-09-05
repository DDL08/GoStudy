package utils

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"reflect"
	"runtime"
	"time"
	"github.com/DDL08/GoStudy/v2/crypto"
)

/* -------------------------- encoding/decoding ------------------------------ */

// UUIDToArray32 UUID to 32byte arrary
func UUIDToArray32(hashID string) [32]byte {
	currnetHashIDSlice, _ := hex.DecodeString(hashID)
	return SliceTo32ByteArray(currnetHashIDSlice)
}

// Array32ToUUID 32byte array to UUID
func Array32ToUUID(bytes [32]byte) string {
	return fmt.Sprintf("%x", bytes)
}

/* -----------------------  Random number --------------------------- */

// NewUUID New Universally Unique Identifier
// TODO use https://github.com/satori/go.uuid to generate UUID
func NewUUID() string {
	return crypto.Sha256(time.Now().Format(time.ANSIC) + GetRandomString(5))
}

// GetRandomString generate a random string
func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

/* -----------------------  Type conversion字节转无符号数字 --------------------------- */

func Uint32ToBytes(i uint32) []byte {
	var buf = make([]byte, 4)
	//把 uint32（4 字节）转为长度为 4 的 []byte。
	binary.BigEndian.PutUint32(buf, uint32(i))
	//使用 大端序（BigEndian）：高位在前，符合网络字节序标准。
	return buf
	/*举例:->
	uint32(i) 将 i 转换为一个 uint32 类型的整数。假设 i 是 12345678，那么 uint32(i) 的值为 12345678。
	大端字节序存储过程：
		12345678 的二进制表示为 00000000 10111100 00111010 00101110（32 位）。
		第一个字节：00000000（十进制 0）
		第二个字节：10111100（十进制 188）
		第三个字节：00111010（十进制 58）
		第四个字节：00101110（十进制 46）
		buf 的最终值为 [0, 188, 58, 46]
	*/
}

// uint32 的最大值是4294967295
func BytesToUint32(buf []byte) uint32 {
	return uint32(binary.BigEndian.Uint32(buf))
}

func Uint16ToBytes(i uint16) []byte {
	var buf = make([]byte, 2)
	binary.BigEndian.PutUint16(buf, uint16(i))
	return buf
}

func BytesToUint16(buf []byte) uint16 {
	return uint16(binary.BigEndian.Uint16(buf))
}

func Uint64ToBytes(i uint64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToUint64(buf []byte) uint64 {
	return uint64(binary.BigEndian.Uint64(buf))
}

/* ----------------------- 👇 --------------------------- */

func SliceTo32ByteArray(buf []byte) [32]byte {
	var result [32]byte // 声明一个长度为32的字节数组
	for i := 0; i < 32; i++ {
		result[i] = buf[i] // 把 buf 的前32个字节复制进去
	}
	return result
}

//常用于哈希（如 SHA256）、UUID 等需要固定长度字节数组的场景。

func IpToUint32(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

// IPv4 转
func Uint32ToIp(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
	// 创建一个长度为4的字节切片，表示IPv4地址的4个字节
	// 把nn这个32位无符号整数以大端字节序写入ip切片中
}

/* -----------------------  system info --------------------------- */
func GetSystemType() (osType uint32) {
	var os = runtime.GOOS
	//Go 的标准库变量，表示当前编译环境的操作系统
	switch os {
	case "darwin":
		osType = 0x00
	case "windows":
		osType = 0x01
	case "linux":
		osType = 0x02
	default:
		// unknown
		osType = 0xff
	}
	return
}

/* ------------------------ file operation ---------------------------- */

// FileExists Check if the file exists
func FileExists(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// IsDir check if path is dir
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// GetFileSize Get the size of a single file
func GetFileSize(path string) int64 {
	fileInfo, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	fileSize := fileInfo.Size() //获取size
	return fileSize
}

// GetFileSizeDescription 10.00G / 100.00M
func GetFileSizeDescription(fileSize uint64) string {
	fileSizeFloat := float64(fileSize)
	const K = 1024 * 1.0
	const M = K * 1024
	const G = M * 1024
	if fileSizeFloat >= G {
		return fmt.Sprintf("%0.2fG", fileSizeFloat/G)
	}
	if fileSizeFloat >= M {
		return fmt.Sprintf("%0.2fM", fileSizeFloat/M)
	}
	return fmt.Sprintf("%0.2fK", fileSizeFloat/K)
}

/* -------------- Calculate structure size计算结构体的大小 ----------------- */
func PacketSize(packet interface{}) (uint64, error) {

	var size uint64

	size = 0

	t := reflect.TypeOf(packet)
	v := reflect.ValueOf(packet)

	if k := t.Kind(); k != reflect.Struct {
		return 0, errors.New("Param is Not Struct")
	}

	count := t.NumField()
	for i := 0; i < count; i++ {
		val := v.Field(i).Interface()

		// type switch
		switch value := val.(type) {
		case uint16:
			size += 2
		case uint32:
			size += 4
		case uint64:
			size += 8
		case string:
			size += uint64(len(value))
		case []byte:
			size += uint64(len(value))
		case [2]byte:
			size += uint64(len(value))
		case [4]byte:
			size += uint64(len(value))
		case [32]byte:
			size += uint64(len(value))
		default:
			log.Fatalln("[-]PacketSize error: type unsupport")
			return 0, errors.New("Type unsupport")
		}
	}
	return size, nil
}

/* ------------------- slice deduplication数组去重 ---------------------*/
func RemoveDuplicateElement(addrs []string) []string {
	result := make([]string, 0, len(addrs))
	temp := map[string]struct{}{}
	for _, item := range addrs {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

/* --------------------- handling windows \r ------------------------------*/
func HandleWindowsCR() {
	if runtime.GOOS == "windows" {
		var noUse string
		fmt.Scanf("%s", &noUse)
		// fmt.Println("windows :", []byte(noUse))
	}
}
