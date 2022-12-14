package algoI_test

import (
	"masteralgo/internal/core/domain"
	"masteralgo/leetcode/algoI"
	"masteralgo/pkg/helpers"
	"reflect"
	"testing"
)

type (
	arrayTest struct {
		arg1           []int
		arg2, expected int
	}

	arrayTestV2 struct {
		arg1     []int
		expected []int
	}

	arrayTestV3 struct {
		arg1     []int
		arg2     int
		expected []int
	}

	arrayTestV4 struct {
		arg1             [][]int
		arg2, arg3, arg4 int
		expected         [][]int
	}

	arrayTestV5 struct {
		arg1     [][]int
		expected int
	}

	arrayTestV6 struct {
		arg1     [][]int
		expected [][]int
	}

	arrayTestV7 struct {
		arg1, arg2, expected []int
	}

	arrayTestV8 struct {
		arg1     []int
		expected [][]int
	}

	arrayTestV9 struct {
		arg1, arg2 int
		expected   [][]int
	}

	byteArrayTest struct {
		arg1     []byte
		expected []byte
	}

	badVersionTest struct {
		nbVer, badVer, expected int
	}

	stringTest struct {
		arg1, expected string
	}

	stringTestV2 struct {
		arg1     string
		expected int
	}

	stringTestV3 struct {
		arg1, arg2 string
		expected   bool
	}

	stringTestV4 struct {
		arg1     string
		expected []string
	}

	treeTest struct {
		arg1, arg2, expected []interface{}
	}

	intTest struct {
		arg1     uint32
		expected int
	}

	intTestV2 struct {
		arg1, expected int
	}
)

var testcase704 = []arrayTest{
	{[]int{-1, 0, 3, 5, 9, 12}, -1, 0},
	{[]int{-1, 0, 3, 5, 9, 12}, 0, 1},
	{[]int{-1, 0, 3, 5, 9, 12}, 3, 2},
	{[]int{-1, 0, 3, 5, 9, 12}, 5, 3},
	{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
	{[]int{-1, 0, 3, 5, 9, 12}, 12, 5},
	{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
	{[]int{-1, 0, 3, 5, 9, 12}, 69, -1},
	{[]int{0}, 69, -1},
	{[]int{0}, 0, 0},
	{[]int{0, 1}, 2, -1},
	{[]int{0, 1}, 1, 1},
}

func Test704(t *testing.T) {
	for idx, test := range testcase704 {
		if output := algoI.BinarySearch(test.arg1, test.arg2); output != test.expected {
			t.Errorf("TEST ID: %d. Expected %d but got %d", idx, test.expected, output)
		}
	}
}

func Test704V2(t *testing.T) {
	for idx, test := range testcase704 {
		if output := algoI.BinarySearchV2(test.arg1, test.arg2); output != test.expected {
			t.Errorf("TEST ID: %d. Expected %d but got %d", idx, test.expected, output)
		}
	}
}

var testcase278 = []badVersionTest{
	{5, 4, 4},
	{1, 1, 1},
}

func Test278(t *testing.T) {
	for idx, test := range testcase278 {
		algoI.SetBadVersion(test.badVer)
		if output := algoI.FirstBadVersion(test.nbVer); output != test.expected {
			t.Errorf("TEST ID: %d. Expected %d but got %d", idx, test.expected, output)
		}
	}
}

var testcase35 = []arrayTest{
	{[]int{1, 3, 5, 6}, 1, 0},
	{[]int{1, 3, 5, 6}, 3, 1},
	{[]int{1, 3, 5, 6}, 5, 2},
	{[]int{1, 3, 5, 6}, 6, 3},
	{[]int{1, 3, 5, 6}, 7, 4},
	{[]int{1, 3, 5, 6}, 2, 1},
	{[]int{1, 3}, 2, 1},
	{[]int{0}, 0, 0},
	{[]int{0}, -1, 0},
}

func Test35(t *testing.T) {
	for idx, test := range testcase35 {
		if output := algoI.SearchInsert(test.arg1, test.arg2); output != test.expected {
			t.Errorf("TEST ID: %d. Expected %d but got %d", idx, test.expected, output)
		}
	}
}

var testcase977 = []arrayTestV2{
	{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
	{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
	{[]int{0, 2}, []int{0, 4}},
	{[]int{-1, 2, 2}, []int{1, 4, 4}},
	{[]int{-5, -3, -2, -1}, []int{1, 4, 9, 25}},
	{[]int{1, 2, 3, 4}, []int{1, 4, 9, 16}},
}

func Test977(t *testing.T) {
	for idx, test := range testcase977 {
		output := algoI.SortedSquares(test.arg1)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, output)
		}
	}
}

func Test977V2(t *testing.T) {
	for idx, test := range testcase977 {
		output := algoI.SortedSquaresV2(test.arg1)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, output)
		}
	}
}

var testcase189 = []arrayTestV3{
	{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
	{[]int{1, 2, 3, 4, 5, 6}, 3, []int{4, 5, 6, 1, 2, 3}},
	{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
	{[]int{-1, -100, 3, 99, 2}, 2, []int{99, 2, -1, -100, 3}},
	{[]int{-1, -100, 3}, 2, []int{-100, 3, -1}},
	{[]int{-1}, 2, []int{-1}},
	{[]int{-1, -100, 3, 99, 2, 7}, 2, []int{2, 7, -1, -100, 3, 99}},
}

func Test189(t *testing.T) {
	for idx, test := range testcase189 {
		arg1 := make([]int, len(test.arg1))
		copy(arg1, test.arg1)
		algoI.RotateArray(arg1, test.arg2)
		if !reflect.DeepEqual(arg1, test.expected) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, arg1)
		}
	}
}

func Test189V2(t *testing.T) {
	for idx, test := range testcase189 {
		arg1 := make([]int, len(test.arg1))
		copy(arg1, test.arg1)
		algoI.NaiveRotateArray(arg1, test.arg2)
		if !reflect.DeepEqual(arg1, test.expected) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, arg1)
		}
	}
}

var testcase283 = []arrayTestV2{
	{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
	{[]int{0}, []int{0}},
	{[]int{1, 2, 0, 3}, []int{1, 2, 3, 0}},
	{[]int{1, 0, 0, 0, 3, 2}, []int{1, 3, 2, 0, 0, 0}},
	{[]int{0, 0, 0, 3, 2}, []int{3, 2, 0, 0, 0}},
	{[]int{0, 1, 0, 3, 0, 12}, []int{1, 3, 12, 0, 0, 0}},
}

func Test283(t *testing.T) {
	for idx, test := range testcase283 {
		algoI.MoveZeroes(test.arg1)
		if !reflect.DeepEqual(test.arg1, test.expected) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, test.arg1)
		}
	}
}

var testcase167 = []arrayTestV3{
	{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
	{[]int{2, 3, 4}, 6, []int{1, 3}},
	{[]int{-1, 0}, -1, []int{1, 2}},
}

func Test167(t *testing.T) {
	for idx, test := range testcase167 {
		output := algoI.TwoSumII(test.arg1, test.arg2)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, output)
		}
	}
}

func Test167V2(t *testing.T) {
	for idx, test := range testcase167 {
		output := algoI.TwoSumIIV2(test.arg1, test.arg2)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, output)
		}
	}
}

var testcase344 = []byteArrayTest{
	{[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
	{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}, []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
}

func Test344(t *testing.T) {
	for idx, test := range testcase344 {
		algoI.ReverseString(test.arg1)
		if !reflect.DeepEqual(test.arg1, test.expected) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, test.arg1)
		}
	}
}

var testcase557 = []stringTest{
	{"Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
	{"God Ding", "doG gniD"},
}

func Test557(t *testing.T) {
	for idx, test := range testcase557 {
		output := algoI.ReverseWordsIII(test.arg1)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, output)
		}
	}
}

var testcase876 = []arrayTestV2{
	{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5}},
	{[]int{1, 2, 3, 4, 5, 6}, []int{4, 5, 6}},
}

func Test876(t *testing.T) {
	for idx, test := range testcase876 {
		head := helpers.GetLinkedListfromIntArray(test.arg1)
		middleN := algoI.MiddleNode(head)
		output := helpers.GetIntArrayfromLinkedList(middleN)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, output)
		}
	}
}

var testcase19 = []arrayTestV3{
	{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
	{[]int{1}, 1, []int{}},
	{[]int{1, 2}, 1, []int{1}},
	{[]int{1, 2}, 2, []int{2}},
}

func Test19(t *testing.T) {
	for idx, test := range testcase19 {
		head := helpers.GetLinkedListfromIntArray(test.arg1)
		head = algoI.RemoveNthFromEnd(head, test.arg2)
		output := helpers.GetIntArrayfromLinkedList(head)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, output)
		}
	}
}

var testcase3 = []stringTestV2{
	{"abcabcbb", 3},
	{"bbbbb", 1},
	{"pwwkew", 3},
	{"", 0},
	{"abcdefdgh", 6},
	{"a1aaaaa", 2},
	{"a1aaaaa234", 4},
	{"racecar", 4},
	{"hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", 55},
}

func Test3(t *testing.T) {
	for idx, test := range testcase3 {
		output := algoI.LengthOfLongestSubstring(test.arg1)
		if test.expected != output {
			t.Errorf("TEST ID: %d. Expected %v but got %v in %s", idx, test.expected, output, test.arg1)
		}
	}
}

func Test3V2(t *testing.T) {
	for idx, test := range testcase3 {
		output := algoI.LengthOfLongestSubstringV2(test.arg1)
		if test.expected != output {
			t.Errorf("TEST ID: %d. Expected %v but got %v in %s", idx, test.expected, output, test.arg1)
		}
	}
}

var testcase567 = []stringTestV3{
	{"adc", "dcda", true},
	{"ab", "eidbaooo", true},
	{"ab", "eidboaoo", false},
	{"abc", "eidboaoo", false},
	{"aoob", "eidboaoo", true},
	{"aoobd", "eidboaoo", true},
	{"aoobi", "eidboaoo", false},
}

func Test567(t *testing.T) {
	for idx, test := range testcase567 {
		output := algoI.CheckInclusion(test.arg1, test.arg2)
		if test.expected != output {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, output)
		}
	}
}

var testcase733 = []arrayTestV4{
	{[][]int{
		{1, 1, 1}, {1, 1, 0}, {1, 0, 1},
	}, 1, 1, 2, [][]int{
		{2, 2, 2}, {2, 2, 0}, {2, 0, 1},
	}},
	{[][]int{
		{0, 0, 0}, {0, 0, 0},
	}, 0, 0, 0, [][]int{
		{0, 0, 0}, {0, 0, 0},
	}},
}

func Test733(t *testing.T) {
	for idx, test := range testcase733 {
		output := algoI.FloodFill(test.arg1, test.arg2, test.arg3, test.arg4)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, output)
		}
	}
}

var testcase695 = []arrayTestV5{
	{[][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	},
		6,
	},
	{[][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
	},
		0,
	},
	{[][]int{
		{0, 1}, {1, 0},
	},
		1,
	},
}

func Test695(t *testing.T) {
	for idx, test := range testcase695 {
		output := algoI.MaxAreaOfIsland(test.arg1)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, output)
		}
	}
}

var testcase617 = []treeTest{
	{[]interface{}{}, []interface{}{}, []interface{}{}},
	{[]interface{}{1}, []interface{}{1, 2}, []interface{}{2, 2}},
	{[]interface{}{1, nil, 1, nil, 1}, []interface{}{1, 2}, []interface{}{2, 2, 1, nil, nil, nil, 1}},
	{[]interface{}{1, nil, 1, nil, 1, nil, 1, nil, 1, nil, 1, nil, 1, nil, 1, nil, 1, nil, 1, nil, 1, 2}, []interface{}{1, nil, 1, nil, 1, nil, 1, nil, 1, nil, 1, 2}, []interface{}{2, nil, 2, nil, 2, nil, 2, nil, 2, nil, 2, 2, 1, nil, nil, nil, 1, nil, 1, nil, 1, nil, 1, 2}},
	{[]interface{}{1, 2, nil, 3}, []interface{}{1, nil, 2, nil, 3}, []interface{}{2, 2, 2, 3, nil, nil, 3}},
	{[]interface{}{1, 3, 2, 5}, []interface{}{2, 1, 3, nil, 4, nil, 7}, []interface{}{3, 4, 5, 5, 4, nil, 7}},
}

func Test617(t *testing.T) {
	var err error
	for idx, test := range testcase617 {
		var input1 *domain.TreeNode
		if len(test.arg1) > 0 {
			var root1 domain.ITreeNode
			root1, err = input1.GetTreeFromList(test.arg1)
			if err != nil {
				t.Errorf("TEST ID: %d. Error when parsing first arg: %v", idx, err)
				continue
			}
			input1 = root1.(*domain.TreeNode)
		}
		var input2 *domain.TreeNode
		if len(test.arg2) > 0 {
			var root2 domain.ITreeNode
			root2, err = input2.GetTreeFromList(test.arg2)
			if err != nil {
				t.Errorf("TEST ID: %d. Error when parsing second arg: %v", idx, err)
				continue
			}
			input2 = root2.(*domain.TreeNode)
		}
		output := algoI.MergeTrees(input1, input2)
		outputParsed := output.GetListFromTree()
		if len(test.expected) != 0 || len(outputParsed) != 0 {
			if !reflect.DeepEqual(outputParsed, test.expected) {
				t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, outputParsed)
			}
		}
	}
}

var testcase542 = []arrayTestV6{
	{
		[][]int{
			{1, 1, 1, 1, 1, 1, 0, 1, 1, 1}, {1, 1, 0, 0, 0, 0, 0, 1, 1, 1}, {0, 1, 0, 1, 1, 1, 1, 0, 0, 0}, {1, 1, 1, 0, 0, 1, 1, 0, 1, 1}, {1, 0, 1, 1, 1, 0, 1, 1, 1, 1}, {1, 1, 0, 0, 1, 0, 1, 1, 1, 1}, {1, 0, 1, 0, 0, 0, 1, 1, 0, 1}, {1, 1, 0, 1, 1, 1, 1, 0, 0, 1}, {1, 1, 1, 1, 0, 0, 0, 1, 1, 0}, {1, 1, 1, 0, 1, 1, 0, 1, 1, 1},
		},
		[][]int{
			{2, 2, 1, 1, 1, 1, 0, 1, 2, 2}, {1, 1, 0, 0, 0, 0, 0, 1, 1, 1}, {0, 1, 0, 1, 1, 1, 1, 0, 0, 0}, {1, 1, 1, 0, 0, 1, 1, 0, 1, 1}, {1, 0, 1, 1, 1, 0, 1, 1, 2, 2}, {2, 1, 0, 0, 1, 0, 1, 2, 1, 2}, {1, 0, 1, 0, 0, 0, 1, 1, 0, 1}, {2, 1, 0, 1, 1, 1, 1, 0, 0, 1}, {3, 2, 1, 1, 0, 0, 0, 1, 1, 0}, {3, 2, 1, 0, 1, 1, 0, 1, 2, 1},
		},
	},
	{
		[][]int{
			{0, 0, 0}, {0, 1, 0}, {1, 1, 1},
		},
		[][]int{
			{0, 0, 0}, {0, 1, 0}, {1, 2, 1},
		},
	},
	{
		[][]int{
			{0, 0, 0}, {0, 1, 0}, {0, 0, 0},
		},
		[][]int{
			{0, 0, 0}, {0, 1, 0}, {0, 0, 0},
		},
	},

	{
		[][]int{
			{1, 0, 0}, {1, 1, 1}, {1, 1, 1},
		},
		[][]int{
			{1, 0, 0}, {2, 1, 1}, {3, 2, 2},
		},
	},
	{
		[][]int{
			{1, 0, 1, 1, 0, 0, 1, 0, 0, 1}, {0, 1, 1, 0, 1, 0, 1, 0, 1, 1}, {0, 0, 1, 0, 1, 0, 0, 1, 0, 0}, {1, 0, 1, 0, 1, 1, 1, 1, 1, 1}, {0, 1, 0, 1, 1, 0, 0, 0, 0, 1}, {0, 0, 1, 0, 1, 1, 1, 0, 1, 0}, {0, 1, 0, 1, 0, 1, 0, 0, 1, 1}, {1, 0, 0, 0, 1, 1, 1, 1, 0, 1}, {1, 1, 1, 1, 1, 1, 1, 0, 1, 0}, {1, 1, 1, 1, 0, 1, 0, 0, 1, 1},
		},
		[][]int{
			{1, 0, 1, 1, 0, 0, 1, 0, 0, 1}, {0, 1, 1, 0, 1, 0, 1, 0, 1, 1}, {0, 0, 1, 0, 1, 0, 0, 1, 0, 0}, {1, 0, 1, 0, 1, 1, 1, 1, 1, 1}, {0, 1, 0, 1, 1, 0, 0, 0, 0, 1}, {0, 0, 1, 0, 1, 1, 1, 0, 1, 0}, {0, 1, 0, 1, 0, 1, 0, 0, 1, 1}, {1, 0, 0, 0, 1, 2, 1, 1, 0, 1}, {2, 1, 1, 1, 1, 2, 1, 0, 1, 0}, {3, 2, 2, 1, 0, 1, 0, 0, 1, 1},
		},
	},
	{
		[][]int{
			{0, 1, 0, 1, 1}, {1, 1, 0, 0, 1}, {0, 0, 0, 1, 0}, {1, 0, 1, 1, 1}, {1, 0, 0, 0, 1},
		},
		[][]int{
			{0, 1, 0, 1, 2}, {1, 1, 0, 0, 1}, {0, 0, 0, 1, 0}, {1, 0, 1, 1, 1}, {1, 0, 0, 0, 1},
		},
	},
}

func Test542(t *testing.T) {
	for idx, test := range testcase542 {
		output := algoI.UpdateMatrix(test.arg1)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("TEST ID: %d. Expected \n%v \nbut got \n%v", idx, test.expected, output)
		}
	}
}

func Test542V2(t *testing.T) {
	for idx, test := range testcase542 {
		output := algoI.UpdateMatrixV2(test.arg1)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("TEST ID: %d. Expected \n%v \nbut got \n%v", idx, test.expected, output)
		}
	}
}

var testcase21 = []arrayTestV7{
	{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
	{[]int{}, []int{}, []int{}},
	{[]int{}, []int{0}, []int{0}},
}

func Test21(t *testing.T) {
	for idx, test := range testcase21 {
		head1 := helpers.GetLinkedListfromIntArray(test.arg1)
		head2 := helpers.GetLinkedListfromIntArray(test.arg2)
		output := algoI.MergeTwoLists(head1, head2)
		parsedOutput := helpers.GetIntArrayfromLinkedList(output)
		if !reflect.DeepEqual(parsedOutput, test.expected) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, parsedOutput)
		}
	}
}

var testcase206 = []arrayTestV2{
	{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
	{[]int{1, 2}, []int{2, 1}},
	{[]int{0}, []int{0}},
	{[]int{}, []int{}},
}

func Test206(t *testing.T) {
	for idx, test := range testcase206 {
		head := helpers.GetLinkedListfromIntArray(test.arg1)
		output := algoI.ReverseList(head)
		outputParsed := helpers.GetIntArrayfromLinkedList(output)
		if !reflect.DeepEqual(outputParsed, test.expected) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, outputParsed)
		}
	}
}

var testcase944 = []arrayTestV5{
	{[][]int{
		{2, 1, 1}, {1, 1, 0}, {0, 1, 1},
	},
		4,
	},
	{[][]int{
		{2, 1, 1}, {0, 1, 1}, {1, 0, 1},
	},
		-1,
	},
	{[][]int{
		{0, 2},
	},
		0,
	},
	// no fresh oranges -> 0 minutes
	{[][]int{
		{0},
	},
		0,
	},
	{[][]int{
		{0, 1},
	},
		-1,
	},
	{[][]int{
		{1},
	},
		-1,
	},
	{[][]int{
		{2},
	},
		0,
	},
}

func Test944(t *testing.T) {
	for idx, test := range testcase944 {
		output := algoI.OrangesRotting(test.arg1)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, output)
		}
	}
}

var testcase46 = []arrayTestV8{
	{
		[]int{1, 2, 3},
		[][]int{
			{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1},
		},
	},
	{
		[]int{1},
		[][]int{
			{1},
		},
	},
	{
		[]int{0, 1},
		[][]int{
			{0, 1}, {1, 0},
		},
	},
}

func Test46(t *testing.T) {
	for idx, test := range testcase46 {
		output := algoI.Permute(test.arg1)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, output)
		}
	}
}

var testcase784 = []stringTestV4{
	{"a1b2", []string{"a1b2", "a1B2", "A1b2", "A1B2"}},
	{"3z4", []string{"3z4", "3Z4"}},
}

func Test784(t *testing.T) {
	for idx, test := range testcase784 {
		output := algoI.LetterCasePermutation(test.arg1)
		if !helpers.IsEqualWithoutOrder(output, test.expected) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, output)
		}
	}
}

var testcase77 = []arrayTestV9{
	{4, 2, [][]int{
		{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4},
	}},
	{1, 1, [][]int{
		{1},
	}},
}

func Test77(t *testing.T) {
	for idx, test := range testcase77 {
		output := algoI.Combine(test.arg1, test.arg2)
		if !helpers.IsEqualWithoutOrder(output, test.expected) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, output)
		}
	}
}

var testcase191 = []intTest{
	{00000000000000000000000000001011, 3},
	{00000000000000000000000010000000, 1},
	// {1111111111111111111111111111101, 31},
	{00000000000000000000000010000000, 1},
}

func Test191(t *testing.T) {
	for idx, test := range testcase191 {
		output := algoI.HammingWeight(test.arg1)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, output)
		}
	}
}

func Test191V2(t *testing.T) {
	for idx, test := range testcase191 {
		output := algoI.HammingWeightV2(test.arg1)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, output)
		}
	}
}

// we don't use arg2, so just leave it as 0 in all test cases
var testcase136 = []arrayTest{
	{[]int{2, 2, 1}, 0, 1},
	{[]int{4, 1, 2, 1, 2}, 0, 4},
	{[]int{1}, 0, 1},
}

func Test136(t *testing.T) {
	for idx, test := range testcase136 {
		if output := algoI.SingleNumber(test.arg1); output != test.expected {
			t.Errorf("TEST ID: %d. Expected %d but got %d", idx, test.expected, output)
		}
	}
}

func Test136V2(t *testing.T) {
	for idx, test := range testcase136 {
		if output := algoI.SingleNumberV2(test.arg1); output != test.expected {
			t.Errorf("TEST ID: %d. Expected %d but got %d", idx, test.expected, output)
		}
	}
}

var testcase70 = []intTestV2{
	{2, 2},
	{3, 3},
}

func Test70(t *testing.T) {
	for idx, test := range testcase70 {
		if output := algoI.ClimbStairs(test.arg1); output != test.expected {
			t.Errorf("TEST ID: %d. Expected %d but got %d", idx, test.expected, output)
		}
	}
}

var testcase120 = []arrayTestV5{
	{[][]int{
		{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3},
	},
		11,
	},
	{[][]int{
		{-10},
	},
		-10,
	},
}

func Test120(t *testing.T) {
	for idx, test := range testcase120 {
		output := algoI.MinimumTotal(test.arg1)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("TEST ID: %d. Expected %v but got %v", idx, test.expected, output)
		}
	}
}

// we don't use arg2, so just leave it as 0 in all test cases
var testcase198 = []arrayTest{
	{[]int{1, 2, 3, 1}, 0, 4},
	{[]int{2, 7, 9, 3, 1}, 0, 12},
	{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 0, 0},
}

func Test198(t *testing.T) {
	for idx, test := range testcase198 {
		if output := algoI.Rob(test.arg1); output != test.expected {
			t.Errorf("TEST ID: %d. Expected %d but got %d", idx, test.expected, output)
		}
	}
}

func Test198V2(t *testing.T) {
	for idx, test := range testcase198 {
		if output := algoI.RobV2(test.arg1); output != test.expected {
			t.Errorf("TEST ID: %d. Expected %d but got %d", idx, test.expected, output)
		}
	}
}

func Test198V3(t *testing.T) {
	for idx, test := range testcase198 {
		if output := algoI.RobV3(test.arg1); output != test.expected {
			t.Errorf("TEST ID: %d. Expected %d but got %d", idx, test.expected, output)
		}
	}
}
