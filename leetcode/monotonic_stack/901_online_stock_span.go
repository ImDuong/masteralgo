package monotonicstack

// similar problems:
// 1130. Minimum Cost Tree From Leaf Values
// 907. Sum of Subarray Minimums
// 901. Online Stock Span
// 856. Score of Parentheses
// 503. Next Greater Element II

// Next Greater Element I
// Largest Rectangle in Histogram
// Trapping Rain Water

type Span struct {
	// price: latest price
	// value: max nb of valid consecutive days
	price, value int
}

type StockSpanner struct {
	spans []Span
}

func Constructor() StockSpanner {
	return StockSpanner{
		spans: make([]Span, 0),
	}
}

func (this *StockSpanner) Next(price int) int {
	rs := 1
	// accumulate the span and shrink the span into single value
	for len(this.spans) > 0 && price >= this.spans[len(this.spans)-1].price {
		rs += this.spans[len(this.spans)-1].value
		this.spans = this.spans[:len(this.spans)-1]
	}
	this.spans = append(this.spans, Span{
		price: price,
		value: rs,
	})
	return rs
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
