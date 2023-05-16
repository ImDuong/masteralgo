package algoII

import (
	"masteralgo/leetcode/algoII"
	"reflect"
	"testing"
)

type (
	arrayTestV1 struct {
		arg1     [][]int
		arg2     int
		expected bool
	}

	arrayTestV2 struct {
		arg1     []int
		arg2     int
		expected int
	}
)

var testcase74 = []arrayTestV1{
	{
		[][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		},
		3,
		true,
	},
	{
		[][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 50},
		},
		11,
		true,
	},
	{
		[][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		},
		0,
		false,
	},
	{
		[][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		},
		13,
		false,
	},
	{
		[][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		},
		23,
		true,
	},
	{
		[][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		},
		21,
		false,
	},
	{
		[][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		},
		60,
		true,
	},
	{
		[][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		},
		61,
		false,
	},
	{
		[][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		},
		30,
		true,
	},
}

func Test74(t *testing.T) {
	for idx, test := range testcase74 {
		output := algoII.Search2DMatrix(test.arg1, test.arg2)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, output)
		}
	}
}

var testcase33 = []arrayTestV2{
	{
		[]int{4, 5, 6, 7, 0, 1, 2},
		0,
		4,
	},
	{
		[]int{4, 5, 6, 7, 0, 1, 2},
		3,
		-1,
	},
	{
		[]int{1},
		0,
		-1,
	},
}

func Test33(t *testing.T) {
	for idx, test := range testcase33 {
		output := algoII.SearchRotatedSortedArray(test.arg1, test.arg2)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, output)
		}
	}
}
