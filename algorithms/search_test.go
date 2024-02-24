package algorithms_test

import (
	"com.github/dukendev/basics/algorithms"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	res := algorithms.LinearSearch(list, 3)
	expect := 2
	if res != expect {
		t.Errorf("expected %d but found %d", expect, res)
	} else {
		t.Log("TestLinearSearch passed")
	}
}

func TestBinarySearch(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	res := algorithms.BinarySearch(list, 3)
	expect := true
	if res != expect {
		t.Errorf("expected %t but found %t", expect, res)
	} else {
		t.Log("TestBinarySearch passed")
	}
}

func TestBinarySearchRec(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	res := algorithms.BinarySearchRec(list, 1)
	expect := true
	if res != expect {
		t.Errorf("expected %t but found %t", expect, res)
	} else {
		t.Log("TestBinarySearch passed")
	}
}
