package fib

// Baseline recursive Fibonacci function
func RFib(n int) int {
	if n == 1 || n == 0 {
		return n
	}

	return RFib(n-1) + RFib(n-2)
}

type CachedFib struct {
	cache []int
}

// CFib is a recursive Fibonacci function that memoizes intermediate
// values in an array.
func CFib(n int) int {
	cache := make([]int, 0, n)
	cache = append(cache, 0, 1)

	f := &CachedFib{cache}

	return f.CFib(n)
}

func (self *CachedFib) CFib(n int) (val int) {
	if len(self.cache) == n {
		// calculate and store
		val = (self.CFib(n-1) + self.CFib(n-2))
		self.cache = append(self.cache, val)
	} else if len(self.cache) > n {
		val = self.cache[n]
	} else {
		val = self.CFib(n-1) + self.CFib(n-2)
	}

	return
}

// UFib is an incremental Fibonacci function where we only cache the
// previous two values necessary for calculating the next one.
func UFib(n int) (val int) {
	if n == 0 || n == 1 {
		return n
	}

	back2, back1 := 0, 1

	for i := 2; i <= n; i++ {
		val = back2 + back1
		back2 = back1
		back1 = val
	}

	return
}
