package algoI

// n versions [1, 2, ..., n]
// each version is based on previous versions -> all versions after a bad version is bad
// -> for a sorting list of version, applying Binary Search is a good option

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

var BadVersion int

func isBadVersion(version int) bool {
	return version >= BadVersion
}

func FirstBadVersion(n int) int {
	lowB := 1
	upB := n
	firstBadVer := -1
	for upB >= lowB {
		jmp := (upB + lowB) / 2
		// if jmp is bad version, go left
		// if jmp is right version, go right
		if !isBadVersion(jmp) {
			lowB = jmp + 1
		} else {
			firstBadVer = jmp
			upB = jmp - 1
		}
	}
	return firstBadVer
}

// customizable version to set bad version for testing purpose
func SetBadVersion(version int) {
	BadVersion = version
}
