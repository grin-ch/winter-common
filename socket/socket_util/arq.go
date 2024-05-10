package socket_util

import "github.com/grin-ch/winter-common/collector"

// Arq auto retransmission request
type Arq interface {
	Push(id int64, pack []byte)
	Seek() ([][]byte, []int64)
	Ack(id int64)
}

type quickARQ struct {
	limit  int // 最大重试次数
	counts collector.TypeMap[int64, *delayCount]
}

func NewARQ(limit int) Arq {
	return &quickARQ{
		limit: limit,
	}
}

func (q *quickARQ) Push(id int64, pack []byte) {
	q.counts.Store(id, &delayCount{
		id:    id,
		pack:  pack,
		times: 0,
		delay: 1,
	})
}

func (q *quickARQ) Seek() ([][]byte, []int64) {
	delays := make([][]byte, 0)
	losses := make([]int64, 0)
	q.counts.Range(func(k int64, v *delayCount) bool {
		if v.delay > 0 {
			v.delay--
			return true
		}

		if v.times >= q.limit {
			losses = append(losses, v.id)
			return true
		}

		delays = append(delays, v.pack)
		v.delay = 1 + v.times*2
		v.times++
		return true
	})
	for _, v := range losses {
		q.counts.Delete(v)
	}
	return delays, losses
}

func (q *quickARQ) Ack(id int64) {
	q.counts.Delete(id)
}

type delayCount struct {
	id   int64
	pack []byte

	delay int
	times int
}
