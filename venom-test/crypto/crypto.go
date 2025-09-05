package crypto

import (
	"github.com/DDL08/GoStudy/v2/global"
)

// InitEncryption generate secret key、protocal separator、protocal feature
func InitEncryption(passwd string) {
	if passwd != "" {
		global.SECRET_KEY = Md5Raw(passwd)
		//使用 MD5 对密码进行哈希，作为对称加密通信的密钥。
		global.PROTOCOL_SEPARATOR = string(Md5Raw(passwd + global.PROTOCOL_SEPARATOR)[:4])
		//把密码和原有分隔符字符串连接后再次进行 MD5，取前 4 字节作为新的“协议分隔符”。×

		global.PROTOCOL_FEATURE = string(Md5Raw(passwd + global.PROTOCOL_FEATURE)[:8])
		//取前 8 字节构造一个特征字符串，类似包头校验。×

	} else {

		OVERHEAD = 0
		//无密码时取消加密
	}
}

/*


 */
