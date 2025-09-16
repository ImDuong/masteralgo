package math_challenges

func Challenge2197(nums []int) []int {
	return replaceNonCoprimes(nums)
}

func replaceNonCoprimes(nums []int) []int {
	stk := []int{}
	for i := range nums {
		stk = append(stk, nums[i])
		for len(stk) > 1 {
			a, b := stk[len(stk)-1], stk[len(stk)-2]
			gcd := GCDIterative(a, b)
			if gcd <= 1 {
				break
			}

			// GCD(a, b) * LCM(a, b) = a * b
			lcm := a * b / gcd
			stk = stk[:len(stk)-1]
			stk[len(stk)-1] = lcm
		}
	}
	return stk
}

// GCDIterative Faster iterative version of GcdRecursive without holding up too much of the stack
// time complexity: O(log(min(a, b))) where a and b are the two numbers
// space complexity: O(1)
// Ref: https://github.com/TheAlgorithms/Go/blob/master/math/gcd/gcditerative.go
func GCDIterative(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
