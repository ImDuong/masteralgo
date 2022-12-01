package algoI

// all hail the power of bit manipulation

func IsPowerOfTwo(n int) bool {
	return (n != 0) && (n&(n-1)) == 0
}
