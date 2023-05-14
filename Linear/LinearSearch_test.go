package linear

import "testing"

func TestLinearSearchFound(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	target := 3
	result := Search(arr, target)
	if result != 2 {
		t.Errorf("Expected %d, got %d", 2, result)
	}
}

func TestLinearSearchNotFound(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	target := 6
	result := Search(arr, target)
	if result != -1 {
		t.Errorf("Expected %d, got %d", -1, result)
	}
}

func TestLinearSearchEmptyList(t *testing.T) {
	arr := []int{}
	target := 6
	result := Search(arr, target)
	if result != -1 {
		t.Errorf("Expected %d, got %d", -1, result)
	}
}
