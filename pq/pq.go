/*

A priority queue implementation to practice some Go. Mostly like the
Skiena book implementation. Slightly changed to allow reuse of most code
for both sort directions

Weird decision but I return -1, as a sort of "error" for pq.parent()
but then real errors for Peek() etc. The current reasoning is that
parent returns an int array index, for which -1 is obviously an invalid
value. On the other hand Peek() returns any valid integer, so aside
from min/max value, there's no integer value that could double as an
error-like value.

I could just use errors everywhere but I don't have a clear idea yet
what the Go idioms are around errors and the -1 pattern is from Skiena's
book.

*/
package pq

type PriorityQueue interface {
	Insert(int) error
	Take() (int, error)
	Peek() (int, error)
}

type Heap struct {
	heap []int
	n    int
	Comparator
}

type Comparator interface {
	cmp(int, int) bool
}

type (
	Min struct{}
	Max struct{}
)

// InitMinPriorityQueue returns a pointer to a new PriorityQueue with a
// minimum element at the root. It will be of capacity PQ_SIZE.
func InitMinPriorityQueue(capacity int) *Heap {
	heap := make([]int, capacity, capacity)

	return &Heap{heap, 0, &Min{}}
}

// InitMaxPriorityQueue returns a pointer to a new PriorityQueue with a
// maximum element at the root. It will be of capacity PQ_SIZE.
func InitMaxPriorityQueue(capacity int) *Heap {
	heap := make([]int, capacity, capacity)

	return &Heap{heap, 0, &Max{}}
}

// Inserts to the back of the queue and then bubbles up maintaining the
// invariant that parent value is always greater or smaller than that of
// the children (depending on the direction).
func (q *Heap) Insert(x int) error {
	if q.n >= cap(q.heap) {
		return &QueueOverflowError{}
	}

	q.n++
	q.heap[q.n] = x
	q.bubbleUp(q.n)

	return nil
}

// Take returns the minimum/maximum element and removes it from the
// queue.
func (q *Heap) Take() (int, error) {
	if q.n == 0 {
		return -1, &EmptyQueueError{}
	}

	val := q.heap[1]
	q.heap[1] = q.heap[q.n]
	q.n--
	q.bubbleDown(1)

	return val, nil
}

// Peek returns the min/max value without removing the element from the
// queue.
func (q *Heap) Peek() (int, error) {
	if q.n == 0 {
		return -1, &EmptyQueueError{}
	}

	return q.heap[1], nil
}

func (q *Heap) bubbleUp(p int) {
	if parent(p) == -1 {
		return
	}

	if q.cmp(q.heap[parent(p)], q.heap[p]) {
		q.heap[parent(p)], q.heap[p] = q.heap[p], q.heap[parent(p)]

		q.bubbleUp(parent(p))
	}
}

func (q *Heap) bubbleDown(p int) error {
	if q.n <= 0 {
		return &EmptyQueueError{}
	}

	childIndex := child(p)
	minIndex := p

	for i := 0; i <= 1; i++ {
		if childIndex+i <= q.n {
			if q.cmp(q.heap[minIndex], q.heap[childIndex+i]) {
				minIndex = childIndex + i
			}
		}
	}

	if minIndex != p {
		q.heap[minIndex], q.heap[p] = q.heap[p], q.heap[minIndex]
		q.bubbleDown(minIndex)
	}

	return nil
}

func (q *Min) cmp(x, y int) bool {
	return x > y
}

func (q *Max) cmp(x, y int) bool {
	return x < y
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

type EmptyQueueError struct{}

func (self *EmptyQueueError) Error() string {
	return "Queue is empty."
}

type QueueOverflowError struct{}

func (self *QueueOverflowError) Error() string {
	return "Queue is already full."
}
