package queue

// An FIFO queue
type Queue []int

// Push Pushes element into queue
//	e.g. q.push(123)
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// Pops element from queue
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// Returns if the queue is empty or not
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
