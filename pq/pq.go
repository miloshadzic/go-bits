package pq

const PQ_SIZE = 32

type PriorityQueue struct {
	Heap *[PQ_SIZE]int
	N    int
}

func Init() *PriorityQueue {
	var heap [PQ_SIZE]int

	return &PriorityQueue{&heap, 0}
}

func Parent(n int) int {
	if n == 1 {
		return -1
	}

	return n / 2
}

func Child(n int) int {
	return n * 2
}

func (q *PriorityQueue) Insert(x int) {
	if q.N >= PQ_SIZE {
		return
	} else {
		q.N++
		q.Heap[q.N] = x
		q.BubbleUp(q.N)
	}
}

func (q *PriorityQueue) BubbleUp(p int) {
	if Parent(p) == -1 {
		return
	}

	if q.Heap[Parent(p)] > q.Heap[p] {
		q.Heap[Parent(p)], q.Heap[p] = q.Heap[p], q.Heap[Parent(p)]

		q.BubbleUp(Parent(p))
	}
}
