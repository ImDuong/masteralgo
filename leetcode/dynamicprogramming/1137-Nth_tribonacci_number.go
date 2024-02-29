package dynamicprogramming

// T0 = 0, T1 = 1, T2 = 1, and Tn+3 = Tn + Tn+1 + Tn+2 for n >= 0.
// 0 <= n <= 37
func Tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}

	triArr := make([]int, 38)
	for i := range triArr {
		triArr[i] = -1
	}
	triArr[0] = 0
	triArr[1] = 1
	triArr[2] = 1
	return calTribonacci(n, triArr)
}

func calTribonacci(n int, triArr []int) int {
	var first, second, third int
	if triArr[n-1] != -1 {
		first = triArr[n-1]
	} else {
		first = calTribonacci(n-1, triArr)
	}

	if triArr[n-2] != -1 {
		second = triArr[n-2]
	} else {
		second = calTribonacci(n-2, triArr)
	}

	if triArr[n-3] != -1 {
		third = triArr[n-3]
	} else {
		third = calTribonacci(n-3, triArr)
	}
	triArr[n] = first + second + third
	return triArr[n]
}
