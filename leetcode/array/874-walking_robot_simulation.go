package array

import "fmt"

func Challenge874(commands []int, obstacles [][]int) int {
	return robotSim(commands, obstacles)
}

type robot874 struct {
	maxDist int
	curX    int
	curY    int
	// 1, 2, 3, 4: north, east, south, west
	direction int
}

func (r *robot874) execute(cmd int, obstacles [][]int) {
	switch cmd {
	case -1:
		r.direction = r.direction + 1
		if r.direction > 4 {
			r.direction -= 4
		}
	case -2:
		r.direction = (r.direction - 1) % 4
		if r.direction == 0 {
			r.direction = 4
		}
	default:
		r.curX, r.curY = r.getExpectedCoord(cmd, obstacles)
		r.maxDist = max(r.maxDist, r.curX*r.curX+r.curY*r.curY)
	}
}

func (r *robot874) getExpectedCoord(k int, obstacles [][]int) (exX int, exY int) {
	exX = r.curX
	exY = r.curY

	switch r.direction {
	case 1:
		exX = r.curX
		exY = r.curY + k
		isStopped, obX, obY := r.getClosestObstacle(k, obstacles)
		if isStopped {
			exX = obX
			exY = obY - 1
		}
	case 2:
		exX = r.curX + k
		exY = r.curY
		isStopped, obX, obY := r.getClosestObstacle(k, obstacles)
		if isStopped {
			exX = obX - 1
			exY = obY
		}
	case 3:
		exX = r.curX
		exY = r.curY - k
		isStopped, obX, obY := r.getClosestObstacle(k, obstacles)
		if isStopped {
			exX = obX
			exY = obY + 1
		}
	case 4:
		exX = r.curX - k
		exY = r.curY
		isStopped, obX, obY := r.getClosestObstacle(k, obstacles)
		if isStopped {
			exX = obX + 1
			exY = obY
		}

	}

	return
}

func (r *robot874) getClosestObstacle(k int, obstacles [][]int) (isStopped bool, closestObX, closestObY int) {
	nbObstacles := len(obstacles)
	for i := 0; i < nbObstacles; i++ {
		obX := obstacles[i][0]
		obY := obstacles[i][1]
		if !r.isObstacleInSameAxis(obX, obY) {
			continue
		}
		switch r.direction {
		case 1:
			if obY > r.curY && obY <= r.curY+k {
				if !isStopped {
					closestObX = obX
					closestObY = obY
					isStopped = true
				} else if closestObY > obY {
					closestObX = obX
					closestObY = obY
				}
			}
		case 2:
			if obX > r.curX && obX <= r.curX+k {
				if !isStopped {
					closestObX = obX
					closestObY = obY
					isStopped = true
				} else if closestObX > obX {
					closestObX = obX
					closestObY = obY
				}
			}
		case 3:
			if obY < r.curY && obY >= r.curY-k {
				if !isStopped {
					closestObX = obX
					closestObY = obY
					isStopped = true
				} else if closestObY < obY {
					closestObX = obX
					closestObY = obY
				}
			}
		case 4:
			if obX < r.curX && obX >= r.curX-k {
				if !isStopped {
					closestObX = obX
					closestObY = obY
					isStopped = true
				} else if closestObX < obX {
					closestObX = obX
					closestObY = obY
				}
			}
		}
	}
	return
}

func (r *robot874) isObstacleInSameAxis(obX, obY int) bool {
	switch r.direction {
	case 1:
		fallthrough
	case 3:
		return r.curX == obX
	case 2:
		fallthrough
	case 4:
		return r.curY == obY
	}
	return false
}

func robotSim(commands []int, obstacles [][]int) int {
	nbCmds := len(commands)
	robot := &robot874{
		direction: 1,
	}
	for i := 0; i < nbCmds; i++ {
		robot.execute(commands[i], obstacles)
	}
	fmt.Println("Latest Coord", robot.curX, robot.curY)
	return robot.maxDist
}
