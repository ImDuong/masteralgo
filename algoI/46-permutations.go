package algoI

var permutations = [][]int{}

func Permute(nums []int) [][]int {
	permutations = [][]int{}
	visited := make([]bool, len(nums))
	x := make([]int, len(nums))
	tryPermute(0, nums, x, visited)
	return permutations
}

func tryPermute(k int, nums, permut []int, visited []bool) {
	for v := 0; v < len(nums); v++ {
		if !visited[v] {
			permut[k] = nums[v]
			visited[v] = true
			if k == len(nums)-1 {
				added := make([]int, len(nums))
				copy(added, permut)
				permutations = append(permutations, added)
			} else {
				tryPermute(k+1, nums, permut, visited)
			}
			visited[v] = false
		}
	}
}
