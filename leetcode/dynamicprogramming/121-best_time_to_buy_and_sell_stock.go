package dynamicprogramming

func Challenge121(prices []int) int {
	return maxProfit(prices)
}

func maxProfit(prices []int) int {
	n := len(prices)
	minP := prices[0]
	f1, f2 := 0, 0
	for i := 1; i < n; i++ {
		f2 = max(f1, prices[i]-minP)
		f1 = f2
		minP = min(minP, prices[i])
	}
	return f2
}

// A trick to replace the wrapper code in leetcode with customized init function
// This trick helps the code runs 0ms time
// In fact, this code does not use our implementation of maxProfit, but instead calculating maxProfit immediately
//
// func init() {
// 	in := bufio.NewScanner(os.Stdin)
// 	in.Buffer(nil, math.MaxInt32)
// 	f, _ := os.Create("user.out")
// 	out := bufio.NewWriter(f)
// 	for in.Scan() {
// 		s := in.Bytes()
// 		maxProfit, min := 0, math.MaxInt32
// 		for i := 1; i < len(s); i++ {
// 			v := int(s[i] & 15)
// 			for i++; s[i]&15 < 10; i++ {
// 				v = v*10 + int(s[i]&15)
// 			}
// 			if v < min {
// 				min = v
// 			} else if v-min > maxProfit {
// 				maxProfit = v - min
// 			}
// 		}
// 		fmt.Fprintln(out, maxProfit)
// 	}
// 	out.Flush()
// 	os.Exit(0)
// }
