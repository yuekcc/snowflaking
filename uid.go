package snowflaking

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

const (
	_SEQUENCE_MAX      = 99 // 序号最大值
	_SEQUENCE_STR_SIZE = 2  // 序号转为字符串的长度
)

func formatNumber(num int64) string {
	return fmt.Sprintf("%d", num)
}

// IDWorkder 表示一个序列号生成器
//
type IDWorkder struct {
	workerID      string // Worker ID
	sequence      int64  // 序号
	lastTimestamp int64
	sync.Mutex
}

// NewIDWorkder 构建一个新的 IDWorker，worker ID 必须是千位的数值
//
func NewIDWorkder(id int) (*IDWorkder, error) {
	if id < 1000 || id > 9999 {
		return nil, errors.New("worker ID 超出范围")
	}

	worker := &IDWorkder{
		workerID:      formatNumber(int64(id)),
		lastTimestamp: 0,
		sequence:      1,
	}
	return worker, nil
}

// getTimestamp 生成 Timestamp，精确到毫秒
//
func (w *IDWorkder) getTimestamp() int64 {
	return time.Now().UnixNano() / int64(1000000)
}

func (w *IDWorkder) refreshTimestamp(last int64) int64 {
	t := w.getTimestamp()
	for {
		if t <= last {
			t = w.getTimestamp()
		} else {
			break
		}
	}

	return t
}

// Next 生成流水号
//
func (w *IDWorkder) Next() (string, error) {
	w.Lock()
	defer w.Unlock()

	ts := w.getTimestamp()
	for {
		// 同一个时间点内，序号自增
		if ts == w.lastTimestamp {
			w.sequence += 1
		} else {
			w.sequence = 1
		}

		// 序号超出最大值，重置
		if w.sequence > _SEQUENCE_MAX {
			w.sequence = 0
		} else {
			break
		}

		ts = w.refreshTimestamp(ts)
		w.lastTimestamp = ts
	}

	w.lastTimestamp = ts
	id := w.workerID + formatUnixTimestamp(ts) + formatSequence(w.sequence)
	return id, nil
}

// formatSequence 格式化序号，长度不够的，补零
//
func formatSequence(num int64) string {
	numStr := formatNumber(num)
	return strings.Repeat("0", _SEQUENCE_STR_SIZE-len(numStr)) + numStr
}

func formatUnixTimestamp(ts int64) string {
	sec := ts / 1000
	ms := formatNumber(ts - sec*1000)
	return time.Unix(sec, 0).Format("20060102150405") + strings.Repeat("0", 3-len(ms)) + ms
}
