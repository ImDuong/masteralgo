package monotonicstack

func Challenge42(height []int) int {
	return trap(height)
}

// 2 monotonic stacks but just for tracking max from the left and right (not using stack directly but using the same idea)
// because water only stays when left and right is bigger than current
func trap(height []int) int {
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))

	leftMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	rightMax[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	rs := 0
	for i := range height {
		rs += min(leftMax[i], rightMax[i]) - height[i]
	}
	return rs
}

func trapWithMonotonicStack(height []int) int {
	rs := 0

	// decreasing monotonic stack
	stk := []int{}
	// to trap water, there must be left and right boundary, and hollow between these boundaries
	// by using a decreasing monotonic stack, we can fill water layer by layer when finding the biggest left index corresponding to the current right index
	for right := range height {
		for len(stk) > 0 && height[stk[len(stk)-1]] < height[right] {
			// mid is the hollow we're looking
			mid := stk[len(stk)-1]
			stk = stk[:len(stk)-1]

			// no water can be trapped, if there's no left boundary
			if len(stk) == 0 {
				break
			}

			left := stk[len(stk)-1]

			// water = width * height
			rs += (right - left - 1) * min(height[right]-height[mid], height[left]-height[mid])
		}
		stk = append(stk, right)
	}
	return rs
}
