package node

import "github.com/DDL08/GoStudy/v2/utils"

// CurrentNode 当前节点，这有个Node，对这个结构体进行了一个变量初始化
var CurrentNode = Node{
	IsAdmin: 0,
	HashID:  utils.NewUUID(),
	//生成唯一标识符

	// Conn:          nil, // 当前节点无需使用Conn
	// ConnReadLock:  nil,
	// ConnWriteLock: nil,
	// Socks5SessionIDLock: nil,
}

var GNetworkTopology = NetworkTopology{
	RouteTable: make(map[string]string),		// 表示路由表：目标节点 => 下一跳节点
	NetworkMap: make(map[string]([]string)),	// 表示网络连接图：一个节点 => 与之直接相连的节点列表
}

var GNodeInfo = NodeInfo{
	NodeNumber2UUID: make(map[int]string),
	NodeUUID2Number: make(map[string]int),
	NodeDescription: make(map[string]string),
/*
节点编号 => UUID
UUID => 节点编号
UUID => 描述信息
*/
}

// Nodes 与当前节点连接的节点
var Nodes = make(map[string]*Node)
