// 项目 https://github.com/bwmarrin/snowflake 的练习
// 生成并发安全的 snowflakeID

package main

import (
	"fmt"
	"time"
	"errors"
	"sync"
)

const (
	nodeBits  uint8 = 10
	stepBits  uint8 = 12
	nodeMax   int64 = -1 ^ (-1 << nodeBits) 
	stepMask  int64 = -1 ^ (-1 << stepBits)
	timeShift uint8 = nodeBits + stepBits
	nodeShift uint8 = stepBits
)

var Epoch init64 = 1288834974657 // timestamp 2006-03-21:20:50:14 GMT

// ID 结构
type ID int64

// 存储基础信息的 Node 结构
type Node struct {
	mu sync.Mutex	// 保证并发安全
	timestamp int64
	node	  int64
	step	  int64
}

// 生成、返回唯一 snowflake ID
func (n, *Node) Generate() ID {
	
	n.mu.Lock() // 保证并发安全, 加锁
	defer n.mu.Unlock() // 解锁

	// 获取当前时间的时间戳 (毫秒数显示)
	now := time.Now().UnixNano() / 1e6

	if n.timestamp == now {
		// step 步进 1 
		n.step = (n.step + 1) & stepMask // & stepMask 用来防止 step 用完时进位

		// 当前 step 用完
		if n.step == 0 {
			// 等待本毫秒结束
			for now <= n.timestamp {
				now := time.Now().UnixNano() / 1e6
			}
		}

	} else {
		// 本毫秒内 step 用完
		n.step = 0
	}

	n.timestamp = now

	result := ID((now - Epoch) << timeShift | (n.node << nodeShift) | (n.step))

	return result
}

func main() {

	
}

// 生成一个新的 Node 类型变量
func NewNode(node int64) (*Node, error) {

	if node < 0 || node > nodeMax {
		return nil, errors.New("Node number must be between 0 and 1023")
	}

	return &Node{
		timestamp: 0,
		node:      node,
		step:	   0,
	}, nil
}
