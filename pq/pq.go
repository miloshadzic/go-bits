package pq

const PQ_SIZE = 32

type PriorityQueue struct {
	heap *[PQ_SIZE]int
	n    int
}

func Init() *PriorityQueue {
	var heap [PQ_SIZE]int

	return &PriorityQueue{&heap, 0}
}

// Inserts to the back of the queue and then bubbles up maintaining the invariant that parent value is always greater than that of the children.
func (q *PriorityQueue) Insert(x int) error {
	if q.n >= PQ_SIZE {
		return &QueueOverflowError{}
	}

	q.n++
	q.heap[q.n] = x
	q.bubbleUp(q.n)

	return nil
}

func (q *PriorityQueue) Peek() (int, error) {
	if q.n == 0 {
		return -1, &EmptyQueueError{}
	}

	return q.heap[1], nil
}

// parent returns either a parent index (n / 2) or -1 as an "error".
func parent(n int) int {
	if n == 1 {
		return -1
	}

	return n / 2
}

func child(n int) int {
	return n * 2
}

func (q *PriorityQueue) bubbleUp(p int) {
	if parent(p) == -1 {
		return
	}

	if q.heap[parent(p)] > q.heap[p] {
		q.heap[parent(p)], q.heap[p] = q.heap[p], q.heap[parent(p)]

		q.bubbleUp(parent(p))
	}
}

type EmptyQueueError struct{}

func (self *EmptyQueueError) Error() string {
	return "Queue is empty."
}

type QueueOverflowError struct{}

func (self *QueueOverflowError) Error() string {
	return "Queue is already full."
}
