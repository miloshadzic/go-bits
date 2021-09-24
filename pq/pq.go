/*

A priority queue implementation to practice some Go. Mostly like the
Skiena book implementation.

Weird decision but I return -1, as a sort of "error" for pq.parent() but
then real errors for Peek() etc. The current reasoning is that parent
returns an int array index, for which -1 is obviously an invalid value.
On the other hand Peek() returns any valid integer, so aside from
min/max value, there's no integer value that could double as an
error-like value.

I could just use errors everywhere but I don't have a clear idea yet
what the Go idioms are around errors and the -1 pattern is from Skiena's
book.

*/
package pq

const PQ_SIZE = 32

// A min PriorityQueue/Heap.
type PriorityQueue struct {
	heap *[PQ_SIZE]int
	n    int
}

// Init returns a pointer to a new PriorityQueue. It will be of capacity
// PQ_SIZE. This is quite close to Skiena, but honestly I probably
// should use slices and make the size either dynamic or capped at
// initialization time.
func Init() *PriorityQueue {
	var heap [PQ_SIZE]int

	return &PriorityQueue{&heap, 0}
}

// Inserts to the back of the queue and then bubbles up maintaining the
// invariant that parent value is always greater than that of the
// children.
func (q *PriorityQueue) Insert(x int) error {
	if q.n >= PQ_SIZE {
		return &QueueOverflowError{}
	}

	q.n++
	q.heap[q.n] = x
	q.bubbleUp(q.n)

	return nil
}

// Take returns the minimum element and removes it from the queue.
func (q *PriorityQueue) Take() (int, error) {
	if q.n == 0 {
		return -1, &EmptyQueueError{}
	}

	val := q.heap[1]
	q.heap[1] = q.heap[q.n]
	q.n--
	q.bubbleDown(1)

	return val, nil
}

// Peek returns the min value without removing the element from the
// queue.
func (q *PriorityQueue) Peek() (int, error) {
	if q.n == 0 {
		return -1, &EmptyQueueError{}
	}

	return q.heap[1], nil
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

func (q *PriorityQueue) bubbleDown(p int) (int, error) {
	if q.n <= 0 {
		return -1, &EmptyQueueError{}
	}

	childIndex := child(p)
	minIndex := p

	for i := 0; i <= 1; i++ {
		if childIndex+i <= q.n {
			if q.heap[minIndex] > q.heap[childIndex+i] {
				minIndex = childIndex + i
			}
		}
	}

	if minIndex != p {
		q.heap[minIndex], q.heap[p] = q.heap[p], q.heap[minIndex]
		q.bubbleDown(minIndex)
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

type EmptyQueueError struct{}

func (self *EmptyQueueError) Error() string {
	return "Queue is empty."
}

type QueueOverflowError struct{}

func (self *QueueOverflowError) Error() string {
	return "Queue is already full."
}
