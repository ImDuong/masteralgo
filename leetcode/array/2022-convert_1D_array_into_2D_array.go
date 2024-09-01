package array

func Challenge2022(original []int, m int, n int) [][]int {
	return construct2DArray(original, m, n)
}

func construct2DArray(original []int, m int, n int) [][]int {
	lenOriginal := len(original)
	if m*n != lenOriginal {
		return [][]int{}
	}
	newArray := make([][]int, m)
	for i := 0; i < m; i++ {
		newArray[i] = make([]int, n)
	}
	for k := 0; k < lenOriginal; k++ {
		newArray[k/n][k%n] = original[k]
	}
	return newArray
}
