package queue

// const STAUS_EMPTY = 0 // 队列为空
// const STATUS_SLOW = 1 // 队列满了
// const STATUS_FREE = 2 // 队列非空，空余

type QueueSort struct {
	front, tail int
	maxsize     int
	data        []Elem
	length      int
}

// InitQueue 初始化队列
func (q *QueueSort) InitQueue(size int) {
	q.front = 0
	q.tail = 0
	q.maxsize = size
	q.length = 0
	q.data = make([]Elem, 0, q.maxsize)

}

// EnQueue 入队
func (q *QueueSort) EnQueue(e Elem) bool {
	if q.Slow() {
		return false
	}
	// 实现循环队列
	index := (q.tail + 1) % q.maxsize
	q.data[index-1] = e
	q.tail = index
	q.length++

	return true
}

// OutQueue 出队
func (q *QueueSort) OutQueue(e *Elem) bool {
	if q.Empty() {
		return false
	}

	*e = q.data[q.front]
	index := (q.front + 1) % q.maxsize
	q.front = index
	q.length--

}

// Empty 判断队列状态
func (q *QueueSort) Empty() bool {
	return q.front == q.tail
}

// Slow 判断队列是否满了
func (q *QueueSort) Slow() bool {
	// 因为是循环队列，所以采用浪费一个数组元素的方法来判断队列是否满了，也就是队列最大长度永远是maxsize - 1
	index := (q.tail + 1) % q.maxsize
	return index == q.front
}
