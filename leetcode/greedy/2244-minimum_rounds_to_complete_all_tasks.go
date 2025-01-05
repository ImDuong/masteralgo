package greedy

func Challenge2244(tasks []int) int {
	return minimumRounds(tasks)
}

// Applying Frobenius Coin Problem (Chicken McNugget Theorem): When two numbers are coprime (their greatest common divisor, GCD, is 1),
// every integer greater than or equal to a certain minimum can be expressed as a linear combination of these two numbers with non-negative coefficients.
// -> nbTasks can be represented as 2x + 3y with x, y >= 0. Because GCD(2, 3) = 1
// edge case: if nbTasks is 1, there's no x, y satisfied
// Hence the problem now turns to find x, y that 2x + 3y = nbTasks and x + y is minimum
// -> x + (nbTasks - 2x) / 3 is min
// -> 3x + nbTasks - 2x is min
// -> x is min
// -> hence, we will try to find y first by dividing nbTasks with 3 first. This means y = round_down(nbTasks / 3)
// edge case: if nbTasks % 3 == 1, we need to reduce y by 1. This means y = round_down(nbTasks / 3) - 1
func minimumRounds(tasks []int) int {
	d := make(map[int]int)
	for i := range tasks {
		d[tasks[i]]++
	}
	rs := 0
	for i := range d {
		nbTasks := d[i]
		if nbTasks < 2 {
			return -1
		}
		if nbTasks%3 == 1 {
			nb3 := nbTasks/3 - 1
			nb2 := (nbTasks - nb3*3) / 2
			rs += nb3 + nb2
		} else {
			rs += nbTasks/3 + (nbTasks%3)/2
		}
	}
	return rs
}

// Cleaner way to find x, y that 2x + 3y = nbTasks and x + y is minimum
// Like proving above, we need to find y first by representing nbTasks = 3y + a
// If nbTasks = 3y, we need y times
// If nbTasks = 3y + 1 -> nbTasks = 3 * (y - 1) + 2 + 2 => need y - 1 + 2 times (because 2 + 2 will make 2 times of group 2) -> need y + 1 times
// If nbTasks = 3y + 2 -> need y + 1 times
// => we find that for nbTasks, we will need round_up(nbTasks / 3)
// => equivalents to (nbTasks + 3 - 1) / 3
func minimumRoundsWithCeil(tasks []int) int {
	d := make(map[int]int)
	for i := range tasks {
		d[tasks[i]]++
	}
	rs := 0
	for i := range d {
		nbTasks := d[i]
		if nbTasks < 2 {
			return -1
		}
		rs += (nbTasks + 3 - 1) / 3
	}
	return rs
}
