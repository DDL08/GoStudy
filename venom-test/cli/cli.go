package cli

const (
	LISTEN_MODE  = 1
	CONNECT_MODE = 2
		//客户端和服务端的区别
)
/*
const (
	SOCKET_REUSE_METHOD = 1
		//通过套接字选项 SO_REUSEADDR 和 SO_REUSEPORT 实现端口复用
	IPTABLES_METHOD     = 2
		//通过 Linux 防火墙工具 iptables 实现端口复用
)
*/
type Option struct {
	LocalPort  int
	LocalIP    string
	RemoteIP   string
	RemotePort int
	Mode int
/*
	ReusedPort int

		//模式，0表示未设置或错误，1=监听模式，2=连接模式
	PortReuseMethod int
		//端口复用方法，1=SOCKET_REUSE_METHOD，2=IPTABLES_METHOD
*/
	Password string
}

// Args
var Args Option
