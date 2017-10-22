package queue

import "errors"

type Elem int
type Node struct {
	data Elem
	next *Node
}

type QueueLink struct {
	front  *Node // 对头
	tail   *Node // 队尾
	length int
}

func (q *QueueLink) InitQueue() {
	q.front = new(Node)
	q.tail = q.front
	q.length = 0
}

func (q *QueueLink) EnQueue(e Elem) {
	node := new(Node)
	node.data = e
	node.next = nil

	q.tail.next = node
	q.tail = node
	q.length++
}

func (q *QueueLink) OutQueue(e *Elem) error {
	if q.Empty() {
		return errors.New("empty queue.")
	}

	node := q.front.next
	*e = node.data
	q.front.next = node.next
	// 如果弹出的是队尾元素那么队列就空了，这个时候队尾等于队列
	if q.tail == node {
		q.tail = q.front
	}
	q.length--

	return nil
}

func (q *QueueLink) Empty() bool {
	return q.front == q.tail
}

func (q *QueueLink) Length() int {
	return q.length
}

func (q *QueueLink) Destroy() {
	q.front = nil
	q.tail = nil
}
