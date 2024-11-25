package bfs

func Challenge752(deadends []string, target string) int {
	return openLock(deadends, target)
}

func openLock(deadends []string, target string) int {
	// map for forward rotate
	fr := map[byte]byte{
		'0': '1',
		'1': '2',
		'2': '3',
		'3': '4',
		'4': '5',
		'5': '6',
		'6': '7',
		'7': '8',
		'8': '9',
		'9': '0',
	}

	// map for backward rotate
	br := map[byte]byte{
		'0': '9',
		'1': '0',
		'2': '1',
		'3': '2',
		'4': '3',
		'5': '4',
		'6': '5',
		'7': '6',
		'8': '7',
		'9': '8',
	}

	visited := make(map[string]bool)
	for i := range deadends {
		if deadends[i] == "0000" {
			return -1
		}
		visited[deadends[i]] = true
	}

	rs := 0

	q := []string{}
	q = append(q, "0000")
	visited["0000"] = true

	for len(q) != 0 {
		for _ = range q {
			curCom := q[0]
			q = q[1:]

			if curCom == target {
				return rs
			}

			for i := 0; i < 4; i++ {
				// forward rotate
				tmp := []byte(curCom)
				tmp[i] = fr[tmp[i]]
				tmpStr := string(tmp)
				if _, ok := visited[tmpStr]; !ok {
					q = append(q, tmpStr)
					visited[tmpStr] = true
				}

				// backward rotate
				tmp = []byte(curCom)
				tmp[i] = br[tmp[i]]
				tmpStr = string(tmp)
				if _, ok := visited[tmpStr]; !ok {
					q = append(q, tmpStr)
					visited[tmpStr] = true
				}
			}
		}
		rs++
	}
	return -1
}
