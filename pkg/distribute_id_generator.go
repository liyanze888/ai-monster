package pkg

import (
	"context"
	"fmt"
	"hash/fnv"
)

// DistributeIDGenerator TODO
type DistributeIDGenerator func(key string) (id uint8, err error)

// Produce TODO
func (f DistributeIDGenerator) Produce(key string) (id uint8, err error) {
	return f(key)
}

func defaultGenerateDistributedNodeID(identifier string) (nodeID uint8, err error) {
	// 生成 NODE ID
	if identifier == "" {
		fmt.Println("Snowflake node identifier not set. Node id always 1")
		nodeID = 1
	} else {
		nodeID = hash(identifier)
		//fmt.Printf("Snowflake node identifier set. node identifier = %s, node id = %d\n", identifier, nodeID)
		// 抢占 Node ID
		for i := 0; i < 32; i++ {
			if err = preemptNodeID(identifier, nodeID); err == nil {
				break
			}
			nodeID++
		}
	}
	return
}

func preemptNodeID(identifier string, nodeID uint8) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), HeartbeatInterval)
	defer cancel()
	if success, err := RedisClient.SetNX(ctx, getPreemptRedisKey(nodeID), identifier, OccupiedInterval).Result(); err != nil {
		return fmt.Errorf("cannot occupy snowflake node. nodeID=%d, podId=%s, err=%+v", nodeID, identifier, err)
	} else if !success {
		return fmt.Errorf("cannot occupy snowflake node. nodeID=%d, podId=%s. node id in used", nodeID, identifier)
	}
	if success, err := RedisClient.Expire(ctx, getPreemptRedisKey(nodeID), OccupiedInterval).Result(); err != nil {
		return fmt.Errorf("cannot occupy snowflake node. nodeID=%d, podId=%s, err=%+v", nodeID, identifier, err)
	} else if !success {
		return fmt.Errorf("cannot occupy snowflake node. nodeID=%d, podId=%s. should not happen", nodeID, identifier)
	}
	return
}

func hash(s string) uint8 {
	h := fnv.New32a()
	if _, err := h.Write([]byte(s)); err != nil {
		panic("Hash to get snowflake node id failed")
	}
	return uint8(h.Sum32())
}

func getPreemptRedisKey(nodeID uint8) string {
	return fmt.Sprintf(NodeKeyTemplate, nodeID)
}
