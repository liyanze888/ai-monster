package pkg

import (
	"errors"
	"fmt"
	sf "github.com/bwmarrin/snowflake"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	// FailCallbackFunc node id 失效后的回调函数，一般用于更新健康检查重启 pod
	FailCallbackFunc = func() {}
	// RedisClient 默认使用 redis 进行分布式 node id 配置及心跳机制
	// 使用默认 NodeIDGeneratorFunc 和 HeartBeatFunc 时必须设置
	RedisClient redis.UniversalClient
	// NodeIDGeneratorFunc 允许自定义的分布式ID生成器
	NodeIDGeneratorFunc DistributeIDGenerator = defaultGenerateDistributedNodeID
	// HeartBeatFunc 允许自定义的保持占有 node id 的心跳函数
	HeartBeatFunc HeartBeatExecutor = defaultHeartBeat
	// NodeKeyTemplate TODO
	NodeKeyTemplate = "snowflake:node:id:%d"

	// HeartbeatInterval TODO
	HeartbeatInterval = 5 * time.Second
	// OccupiedInterval TODO
	OccupiedInterval = 40 * time.Second
	// FailThreshold TODO
	// 必须小于OccupiedInterval
	FailThreshold = 20 * time.Second

	lastOccupiedTimeMillis int64
)

var (
	// Epoch 设置 Epoch 为 2020/1/1 8:00
	Epoch int64 = 1577836800000
	// NodeBits 256 nodes
	NodeBits uint8 = 8
	// StepBits 16384 per millisecond
	StepBits uint8 = 14
)

// Node TODO
type Node struct {
	node       *sf.Node
	NodeID     uint8
	Identifier string
}

// NewNode TODO
func NewNode(identifier string) (node *Node, nodeID uint8, err error) {
	// 自定义 snowflake 格式
	customSnowflake()

	if nodeID, err = NodeIDGeneratorFunc(identifier); err != nil {
		return
	}
	lastOccupiedTimeMillis = time.Now().UnixMilli()
	// 创建 snowflake node
	var sNode *sf.Node
	if sNode, err = sf.NewNode(int64(nodeID)); err != nil {
		fmt.Printf("Cannot create snowflake node. nodeID=%d, identifer=%s\n", nodeID, identifier)
		return
	}
	fmt.Printf("Success assign snowflake worker node. node_id = [%d], pod_ip = [%s]\n", nodeID, identifier)
	node = &Node{
		node:       sNode,
		NodeID:     nodeID,
		Identifier: identifier,
	}
	go HeartBeatFunc(node)
	return
}

// Generate TODO
func (n *Node) Generate() (sf.ID, error) {
	if n.Identifier != "" && time.Now().UnixMilli()-lastOccupiedTimeMillis > FailThreshold.Milliseconds() {
		FailCallbackFunc()
		fmt.Println("Snowflake node id occupied timeout.")
		return 0, errors.New("snowflake error. wait for restarting")
	}
	return n.node.Generate(), nil
}

func customSnowflake() {
	sf.Epoch = Epoch
	sf.NodeBits = NodeBits
	sf.StepBits = StepBits
}
