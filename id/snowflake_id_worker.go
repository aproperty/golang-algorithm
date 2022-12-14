package id

import (
	"fmt"
	"sync"
	"time"
)

const (
	nodeBits uint8 = 10 // 工作机器 ID
	stepBits uint8 = 12 // 序列号，即毫秒内的计数

	nodeMax int64 = -1 ^ (-1 << nodeBits)
	stepMax int64 = -1 ^ (-1 << stepBits)

	timeShift uint8 = nodeBits + stepBits
	nodeShift uint8 = stepBits
)

var epoch int64 = 1514764800000 // 2018-01-01 00:00:00  起始时间戳 (毫秒数显示)

// 存储基础信息的 snowflake 结构
type snowflake struct {
	mu sync.Mutex

	node int64

	timestamp int64
	step      int64
}

// 生成、返回 唯一 snowflake ID
func (n *snowflake) Generate() int64 {

	n.mu.Lock()
	defer n.mu.Unlock()

	now := time.Now().UnixNano() / 1e6 // 当前时间的时间戳 (毫秒数显示)

	if n.timestamp == now {
		n.step++
		if n.step > stepMax { // 当前 step 用完
			for now <= n.timestamp { // 等待 本毫秒结束
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		n.step = 0 // 本毫秒内 step 用完
	}

	n.timestamp = now

	result := (now-epoch)<<timeShift | (n.node << nodeShift) | (n.step)
	return result
}

type idWorker interface {
	Generate() int64
}

var IdWorker idWorker

//
func NewIdWorker(node int64) {
	if node < 0 || node > nodeMax {
		panic(fmt.Sprintf("snowflake 节点必须在 0-%d 之间", nodeMax))
	}
	snowflakeIns := &snowflake{
		node: node,

		timestamp: 0,
		step:      0,
	}
	IdWorker = snowflakeIns
}
