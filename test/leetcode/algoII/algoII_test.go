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
		output := algoII.SearchMatrix(test.arg1, test.arg2)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, output)
		}
	}
}
