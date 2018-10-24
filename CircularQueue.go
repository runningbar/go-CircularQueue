package circularqueue

import (
	"errors"
	"sync"
)

// CircularQueue 循环队列
type CircularQueue struct {
	mu     sync.Mutex
	size   int           // queue capacity
	counts []interface{} // queue
	head   int           // 队列头指针
	tail   int           // 队列尾指针
}

// New 初始化一个空队列
func New(s int) *CircularQueue {
	cq := &CircularQueue{
		size:   s,
		counts: make([]interface{}, s),
		head:   -1,
		tail:   -1,
	}
	return cq
}

// GetHead 获取队首count
// 不用锁
func (cq *CircularQueue) GetHead() interface{} {
	return cq.counts[cq.head]
}

// IsFull 判断队列是否已满
func (cq *CircularQueue) IsFull() bool {
	if (cq.tail+1)%cq.size == cq.head {
		return true
	}
	return false
}

// IsEmpty 判断队列是否为空
func (cq *CircularQueue) IsEmpty() bool {
	if cq.head == -1 && cq.tail == -1 {
		return true
	}
	return false
}

// EnQueue 入队，需要锁
func (cq *CircularQueue) EnQueue(count interface{}) bool {
	cq.mu.Lock()
	defer cq.mu.Unlock()

	if cq.IsFull() {
		return false
	}

	if cq.head == -1 {
		cq.head = 0
		cq.tail = 0
	} else {
		cq.tail = (cq.tail + 1) % cq.size
	}

	cq.counts[cq.tail] = count
	return true
}

// DeQueue 出队，需要锁
func (cq *CircularQueue) DeQueue() (interface{}, error) {
	cq.mu.Lock()
	defer cq.mu.Unlock()

	if cq.IsEmpty() {
		return nil, errors.New("Queue is Empty")
	}

	count := cq.counts[cq.head]

	if cq.head == cq.tail {
		cq.head = -1
		cq.tail = -1
	} else {
		cq.head = (cq.head + 1) % cq.size
	}
	return count, nil
}
