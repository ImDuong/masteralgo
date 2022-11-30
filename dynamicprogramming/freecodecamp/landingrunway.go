package dp_freecodecamp

func CanStopRecursive(runway []bool, initSpeed, startIdx int) bool {
	// base case for speed == 0
	if initSpeed == 0 {
		return true
	}

	// other base cases
	if startIdx < 0 || startIdx >= len(runway) || !runway[startIdx] {
		return false
	}

	// note: omit case check initSpeed < 0 due to
	// 1. initSpeed can only be adjusted by 1 unit -> always meet 0 first
	// 2. the original input is always guaranteed non negative

	// try all options of S
	speedOptions := []int{initSpeed, initSpeed + 1, initSpeed - 1}
	for i := range speedOptions {
		if CanStopRecursive(runway, speedOptions[i], startIdx+speedOptions[i]) {
			return true
		}
	}
	return false
}

func CanStopIterative(runway []bool, initSpeed, startIdx int) bool {
	// define the possible maximum speed for a runway
	maxSpeed := len(runway)

	// base cases
	if startIdx < 0 || startIdx >= len(runway) || initSpeed > maxSpeed || !runway[startIdx] {
		return false
	}

	// note: omit case check initSpeed < 0 due to
	// the original input is always guaranteed non negative

	// position i: a set of speeds for which we can stop if we start from position i
	mapPositionToValidSpeeds := make(map[int]map[int]bool)

	// init all landing spots: we can stop anywhere with zero speed
	for position := range runway {
		if runway[position] {
			mapPositionToValidSpeeds[position] = make(map[int]bool)

			// add zero speed to set
			mapPositionToValidSpeeds[position][0] = true
		}
	}

	// for a iterative strategy, we go from a small problem to bigger problems
	// we loop from the end to the start of runway
	for position := len(runway) - 1; position >= 0; position-- {
		if !runway[position] {
			continue
		}

		// go over all possible speeds
		for speed := 1; speed <= maxSpeed; speed++ {
			speedOptions := []int{speed, speed + 1, speed - 1}
			for i := range speedOptions {
				// check if position+speedOptions[i] is a possible position
				if _, ok := mapPositionToValidSpeeds[position+speedOptions[i]]; ok {
					// check if speedOptions[i] is one of the possible speed for the position: position+speedOptions[i]
					if _, ok := mapPositionToValidSpeeds[position+speedOptions[i]][speedOptions[i]]; ok {
						mapPositionToValidSpeeds[position][speed] = true
						break
					}
				}
			}
		}
	}

	if _, ok := mapPositionToValidSpeeds[startIdx]; ok {
		if _, ok := mapPositionToValidSpeeds[startIdx][initSpeed]; ok {
			return true
		}
	}
	return false
}

func CanStopRecursiveWithMemo(runway []bool, initSpeed, startIdx int) bool {
	mapPositionToValidSpeeds := make(map[int]map[int]bool)

	// create a private function to maintain original function signature
	return canStopRecursiveWithMemo(runway, initSpeed, startIdx, mapPositionToValidSpeeds)
}

func canStopRecursiveWithMemo(runway []bool, initSpeed, startIdx int, mapPositionToValidSpeeds map[int]map[int]bool) bool {
	// init the map when needed
	if _, ok := mapPositionToValidSpeeds[startIdx]; !ok {
		mapPositionToValidSpeeds[startIdx] = make(map[int]bool)
	}

	// base case for speed == 0
	if initSpeed == 0 {
		mapPositionToValidSpeeds[startIdx][initSpeed] = true
		return true
	}

	// other base cases
	if startIdx < 0 || startIdx >= len(runway) || !runway[startIdx] {
		mapPositionToValidSpeeds[startIdx][initSpeed] = false
		return false
	}

	// note: omit case check initSpeed < 0 due to
	// 1. initSpeed can only be adjusted by 1 unit -> always meet 0 first
	// 2. the original input is always guaranteed non negative

	// try all options of S
	speedOptions := []int{initSpeed, initSpeed + 1, initSpeed - 1}
	for i := range speedOptions {
		if CanStopRecursive(runway, speedOptions[i], startIdx+speedOptions[i]) {
			mapPositionToValidSpeeds[startIdx][initSpeed] = true
			return true
		}
	}
	mapPositionToValidSpeeds[startIdx][initSpeed] = false
	return false
}
