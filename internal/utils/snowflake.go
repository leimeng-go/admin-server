package utils

import (
	"sync"
	"time"
)

const (
	// 开始时间戳 (2024-01-01 00:00:00)
	startTime = 1704067200000
	// 机器ID所占位数
	workerIDBits = 10
	// 序列号所占位数
	sequenceBits = 12
	// 机器ID最大值
	maxWorkerID = -1 ^ (-1 << workerIDBits)
	// 序列号掩码
	sequenceMask = -1 ^ (-1 << sequenceBits)
	// 机器ID左移位数
	workerIDShift = sequenceBits
	// 时间戳左移位数
	timestampShift = workerIDBits + sequenceBits
)

// Snowflake 雪花ID生成器
type Snowflake struct {
	sync.Mutex
	timestamp int64
	workerID  int64
	sequence  int64
}

// NewSnowflake 创建雪花ID生成器
func NewSnowflake(workerID int64) (*Snowflake, error) {
	if workerID < 0 || workerID > maxWorkerID {
		return nil, nil
	}
	return &Snowflake{workerID: workerID}, nil
}

// NextID 生成下一个ID
func (s *Snowflake) NextID() int64 {
	s.Lock()
	defer s.Unlock()

	now := time.Now().UnixMilli()
	if s.timestamp == now {
		s.sequence = (s.sequence + 1) & sequenceMask
		if s.sequence == 0 {
			// 当前毫秒内计数满了，等待下一毫秒
			for now <= s.timestamp {
				now = time.Now().UnixMilli()
			}
			s.timestamp = now
			s.sequence = 0
		}
	} else {
		s.timestamp = now
		s.sequence = 0
	}

	// 组合ID
	id := ((s.timestamp - startTime) << timestampShift) |
		(s.workerID << workerIDShift) |
		s.sequence

	return id
}
