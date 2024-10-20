package stack

func Challenge1106(expression string) bool {
	return parseBoolExpr(expression)
}

func parseBoolExpr(expression string) bool {
	// s1 holds t, f
	// s2 holds operators
	s1, s2 := newByteStack(len(expression)), newByteStack(len(expression))
	for i := range expression {
		switch expression[i] {
		case '!':
			fallthrough
		case '&':
			fallthrough
		case '|':
			{
				s1.push('#')
				s2.push(expression[i])
			}
		case 't':
			fallthrough
		case 'f':
			{
				s1.push(expression[i])
			}
		case ')':
			{
				// get latest operator
				oper := s2.pop()
				firstBool := s1.pop()
				nextBool := s1.pop()
				if oper == '!' {
					// in this case, nextBool must be '#'
					firstBool = eval(firstBool, nextBool, '!')
				} else {
					for nextBool != '#' {
						firstBool = eval(firstBool, nextBool, oper)
						nextBool = s1.pop()
					}
				}
				// push back evaluated expression
				s1.push(firstBool)
			}
		default:
			{
				continue
			}
		}
	}
	return getBoolean(s1.pop())
}

func eval(n, m, oper byte) byte {
	nBool := getBoolean(n)
	mBool := getBoolean(m)
	switch oper {
	case '!':
		nBool = !nBool
	case '&':
		nBool = nBool && mBool
	case '|':
		nBool = nBool || mBool
	}
	return getByteFromBoolean(nBool)
}

func getBoolean(n byte) bool {
	if n == 't' {
		return true
	}
	return false
}

func getByteFromBoolean(n bool) byte {
	if n == true {
		return 't'
	}
	return 'f'
}

type byteStack struct {
	values []byte
	topIdx int
	cp     int
}

func newByteStack(cp int) *byteStack {
	return &byteStack{
		values: make([]byte, cp+1),
		cp:     cp,
	}
}

func (s *byteStack) push(ele byte) {
	s.values[s.topIdx] = ele
	s.topIdx++
}

func (s *byteStack) pop() byte {
	s.topIdx = max(0, s.topIdx-1)
	return s.values[s.topIdx]
}

func (s *byteStack) top() byte {
	return s.values[s.topIdx-1]
}

func (s *byteStack) isEmpty() bool {
	return s.topIdx == 0
}
