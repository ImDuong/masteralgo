package binarysearch

import "slices"

func Challenge2528(stations []int, r int, k int) int64 {
	return maxPower(stations, r, k)
}

// The challenge related to "maximizing the minimum" -> thinking of using Binary Search
// The condition is monotonic: if we can achieve X min power, we can also achieve any Y < X min power -> Binary Search applicable
// The approach is as follows:
// 1. we find covered power of each city, and find the minPower and maxPower among them
// 2. we do binary search on the answer space [minPower, maxPower + k]
// 3. For each mid, we check if we can achieve mid as the min power by adding at most k stations
// 		3.1. If power at city i < mid, we need to add (mid - power[i]) stations at city (i + r) => greedily
// 		         -> The city (i + r) will cover power for range [i, i + 2r]
//
// If we just use direct array to store the power of each city, updating the power after adding stations will be O(n)
// For each city, leading to O(n^2) per mid check
// The time complexity will be O(n^2 * log(maxPower - minPower)) -> TLE
//
// To optimize the update of power after adding stations, we can use the difference array technique
// to represent the power of each city.
// Using difference array, we can update the power of range [l, r] by just updating two positions in the difference array
// The time complexity of updating power after adding stations will be O(1) per city
// Leading to O(n) per mid check
// Overall time complexity will be O(n * log(maxPower - minPower))
func maxPower(stations []int, r int, k int) int64 {
	diff := make([]int64, len(stations)+1)
	for i := range stations {
		left := max(i-r, 0)
		right := min(i+r+1, len(stations))
		diff[left] += int64(stations[i])
		diff[right] -= int64(stations[i])
	}

	// find the minPower and maxPower among all cities, not just the raw power from stations
	minPower, maxPower := int64(stations[0]), int64(stations[0])
	var power int64
	for i := range stations {
		power += diff[i]
		minPower = min(minPower, power)
		maxPower = max(maxPower, power)
	}

	left, right, mid := minPower, maxPower+int64(k), int64(0)
	for left <= right {
		mid = left + (right-left)/2
		addedStations := int64(k)
		cloned := slices.Clone(diff)
		var power int64
		for i := range stations {
			power += cloned[i]
			if power >= mid {
				continue
			}
			needed := mid - power
			if addedStations < needed {
				addedStations = -1
				break
			}

			addedStations -= needed

			// greedy
			placedIdx := min(len(stations), i+2*r+1)
			cloned[placedIdx] -= needed
			// because we dont care about backward so dont need to increase cloned[i], instead just increase power directly
			power = mid
		}
		if addedStations < 0 {
			// DEBUG: check which stations are failing
			// fmt.Println("failed", mid)
			right = mid - 1
		} else {
			// DEBUG: check which stations are passing
			// fmt.Println("Coool bro", mid)
			left = mid + 1
		}
	}
	return left - 1
}
