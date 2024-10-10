package stack

func Challenge962(nums []int) int {
	return maxWidthRamp(nums)
}

func maxWidthRamp(nums []int) int {
	// stack here used as monotonic stack
	s := newStack(len(nums))

	// first pass: maintain list of indices, where nums[i] is maintained descending
	s.push(0)
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[s.top()] {
			s.push(i)
		}
	}

	// second pass: go from the end of the list to find ramps
	rs := 0
	for i := len(nums) - 1; i >= 0; i-- {
		for !s.isEmpty() && nums[i] >= nums[s.top()] {
			rs = max(rs, i-s.top())
			s.pop()
		}
	}
	return rs
}

type stack struct {
	values []int
	topIdx int
	cp     int
}

func newStack(cp int) *stack {
	return &stack{
		values: make([]int, cp+1),
		cp:     cp,
	}
}

func (s *stack) push(ele int) {
	s.values[s.topIdx] = ele
	s.topIdx++
}

func (s *stack) pop() int {
	s.topIdx = max(0, s.topIdx-1)
	return s.values[s.topIdx]
}

func (s *stack) top() int {
	return s.values[s.topIdx-1]
}

func (s *stack) isEmpty() bool {
	return s.topIdx == 0
}
