package binary

import "testing"

func TestBinarySearchFound(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	target := 3
	result := Search(arr, target)
	if result != 2 {
		t.Errorf("Expected %d, got %d", 2, result)
	}
}

func TestBinarySearchNotFound(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	target := 6
	result := Search(arr, target)
	if result != -1 {
		t.Errorf("Expected %d, got %d", -1, result)
	}
}

func TestBinarySearchEmptyList(t *testing.T) {
	arr := []int{}
	target := 6
	result := Search(arr, target)
	if result != -1 {
		t.Errorf("Expected %d, got %d", -1, result)
	}
}
