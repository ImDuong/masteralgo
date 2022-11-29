package helpers

import (
	"masteralgo/internal/core/domain"
	"math"
	"reflect"
)

func IsEqualWithoutOrder(inp1, inp2 interface{}) bool {
	if reflect.TypeOf(inp1).Kind() != reflect.Slice || reflect.TypeOf(inp2).Kind() != reflect.Slice {
		return false
	}
	input1 := reflect.ValueOf(inp1)
	input2 := reflect.ValueOf(inp2)
	if input1.Len() != input2.Len() {
		return false
	}
	for i := 0; i < input1.Len(); i++ {
		found := false
		for j := 0; j < input2.Len(); j++ {
			if reflect.DeepEqual(input1.Index(i).Interface(), input2.Index(j).Interface()) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func Reverse(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

func GetLinkedListfromIntArray(inp []int) *domain.ListNode {
	if len(inp) == 0 {
		return nil
	}
	head := domain.ListNode{}
	travP := &head
	for i := range inp {
		travP.Val = inp[i]
		if i < len(inp)-1 {
			travP.Next = &domain.ListNode{}
		}
		travP = travP.Next
	}
	return &head
}

func GetIntArrayfromLinkedList(head *domain.ListNode) []int {
	out := []int{}
	travP := head
	for travP != nil {
		out = append(out, travP.Val)
		travP = travP.Next
	}
	return out
}

func GetNumberOfNodesFromLinkedList(head *domain.ListNode) int {
	nbNode := 0
	travNode := head
	for travNode != nil {
		nbNode++
		travNode = travNode.Next
	}
	return nbNode
}

// 4-directionally adjacent cells
func GetAdjacentCellIdx(cellIdx, nbRows, nbCols int) []int {
	adjCells := []int{}
	if cellIdx-nbCols >= 0 {
		adjCells = append(adjCells, cellIdx-nbCols)
	}
	if cellIdx+nbCols < nbRows*nbCols {
		adjCells = append(adjCells, cellIdx+nbCols)
	}
	if (cellIdx%nbCols)-1 >= 0 {
		adjCells = append(adjCells, cellIdx-1)
	}
	if (cellIdx%nbCols)+1 < nbCols {
		adjCells = append(adjCells, cellIdx+1)
	}
	return adjCells
}

func GetRowColIndexes(cellIdx, nbCols int) (int, int) {
	return int(math.Floor(float64(cellIdx) / float64(nbCols))), cellIdx % nbCols
}
