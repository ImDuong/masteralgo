package helpers_test

import (
	"masteralgo/pkg/helpers"
	"reflect"
	"testing"
)

type (
	arrayTest struct {
		arg1     []int
		expected []int
	}

	arrayTestV2 struct {
		arg1     []int
		expected int
	}
)

var testcaseGetLinkedListfromIntArray = []arrayTest{
	{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	{[]int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
}

func TestGetLinkedListfromIntArray(t *testing.T) {
	for idx, test := range testcaseGetLinkedListfromIntArray {
		head := helpers.GetLinkedListfromIntArray(test.arg1)
		output := helpers.GetIntArrayfromLinkedList(head)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, output)
		}
	}
}

var testcaseGetNumberOfNodesFromLinkedList = []arrayTestV2{
	{[]int{1, 2, 3, 4, 5}, 5},
	{[]int{1, 2, 3, 4, 5, 6}, 6},
	{[]int{0}, 1},
	{[]int{}, 0},
}

func TestGetNumberOfNodesFromLinkedList(t *testing.T) {
	for idx, test := range testcaseGetNumberOfNodesFromLinkedList {
		head := helpers.GetLinkedListfromIntArray(test.arg1)
		output := helpers.GetNumberOfNodesFromLinkedList(head)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, output)
		}
	}
}
