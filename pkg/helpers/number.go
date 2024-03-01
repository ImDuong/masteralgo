package helpers

import "math"

// https://floating-point-gui.de/errors/comparison/#look-out-for-edge-cases
func IsRoughlyEqual(a, b, epsilon float64) bool {
	absA := math.Abs(a)
	absB := math.Abs(b)
	diff := math.Abs(a - b)

	if a == b {
		return true
	} else if a == 0 || b == 0 || absA+absB < math.SmallestNonzeroFloat64 {
		return diff < epsilon*math.SmallestNonzeroFloat64
	} else {
		return diff/math.Min(absA+absB, math.MaxFloat64) < epsilon
	}
}

func MinInt(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
