package utils

import "testing"
import "reflect"

// reflect.DeepEqual(a, b)

func TestBatchSplitHappyPath(t *testing.T) {
	src := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	expected := [][]int{[]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}, []int{11, 12}}

	result := BatchSplit(src, 5)

	if !reflect.DeepEqual(result, expected) {
		t.Error(
			"Expected", expected,
			"Got", result,
		)
	}
}

func TestBatchSplitEmpty(t *testing.T) {
	src := []int{}
	expected := [][]int{}

	result := BatchSplit(src, 5)

	if !reflect.DeepEqual(result, expected) {
		t.Error(
			"Expected", expected,
			"Got", result,
		)
	}
}

func TestBatchSplitUnitBatch(t *testing.T) {
	src := []int{1, 2, 3, 4}
	expected := [][]int{[]int{1}, []int{2}, []int{3}, []int{4}}

	result := BatchSplit(src, 1)

	if !reflect.DeepEqual(result, expected) {
		t.Error(
			"Expected", expected,
			"Got", result,
		)
	}
}

func TestBatchSplitSingleBatch(t *testing.T) {
	src := []int{1, 2, 3, 4}
	expected := [][]int{[]int{1, 2, 3, 4}}

	result := BatchSplit(src, 50)

	if !reflect.DeepEqual(result, expected) {
		t.Error(
			"Expected", expected,
			"Got", result,
		)
	}
}

func TestReverseMapHappyPath(t *testing.T) {
	src := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	expected := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}

	result := ReverseMap(src)

	if !reflect.DeepEqual(result, expected) {
		t.Error(
			"Expected", expected,
			"Got", result,
		)
	}
}

func TestFilterHappyPath(t *testing.T) {
	src := []string{"Facebook", "Apple", "Amazon", "Netflix", "Google"}
	stopWords := []string{"Amazon", "Google"}

	expected := []string{"Facebook", "Apple", "Netflix"}

	result := Filter(src, stopWords)

	if !reflect.DeepEqual(result, expected) {
		t.Error(
			"Expected", expected,
			"Got", result,
		)
	}
}

func TestFilterNoStopWords(t *testing.T) {
	src := []string{"Facebook", "Apple", "Amazon", "Netflix", "Google"}
	stopWords := []string{}

	expected := make([]string, len(src))
	copy(expected, src)

	result := Filter(src, stopWords)

	if !reflect.DeepEqual(result, expected) {
		t.Error(
			"Expected", expected,
			"Got", result,
		)
	}
}

func TestFilterNoDeleted(t *testing.T) {
	src := []string{"Facebook", "Apple", "Amazon", "Netflix", "Google"}
	stopWords := []string{"Orange"}

	expected := make([]string, len(src))
	copy(expected, src)

	result := Filter(src, stopWords)

	if !reflect.DeepEqual(result, expected) {
		t.Error(
			"Expected", expected,
			"Got", result,
		)
	}
}
