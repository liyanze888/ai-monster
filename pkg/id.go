package pkg

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/spf13/viper"
	"log"
)

var node *Node

// Generate TODO
func Generate() (snowflake.ID, error) {
	return node.Generate()
}

// GenerateInt64 TODO
func GenerateInt64() (idInt64 int64, err error) {
	var id snowflake.ID
	if id, err = Generate(); err != nil {
		return 0, err
	}
	return id.Int64(), nil
}

// RegisterFailureCallback TODO
func RegisterFailureCallback(f func()) {
	FailCallbackFunc = f
}

func init() {
	var err error
	podIP := viper.GetString("POD_IP")
	if podIP == "" {
		log.Println("POD_IP not set. Snowflake ID unsafe.")
	}
	Epoch = 1577836800000 // 设置 Epoch 为 2020/1/1 8:00
	NodeBits = 8          // 256 nodes
	StepBits = 14         // 16384 per millisecond
	RedisClient = GetRedisClient()
	NodeKeyTemplate = "jojoy:snowflake:node:id:%d"
	node, _, err = NewNode(podIP)
	if err != nil {
		panic(fmt.Sprintf("Cannot init snowflake_config node. podIP=%s, err=%+v", podIP, err))
	}
}
