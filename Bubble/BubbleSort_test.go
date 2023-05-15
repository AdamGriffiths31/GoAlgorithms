package bubble

import "testing"

func TestBubbleSort(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	Sort(arr)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1] < arr[i] {
			t.Errorf("Expected %d, got %d", arr[i], arr[i+1])
		}
	}
}

func TestBubbleSortEmpty(t *testing.T) {
	arr := []int{}
	Sort(arr)
	if len(arr) != 0 {
		t.Errorf("Expected %d, got %d", 0, len(arr))
	}
}

func TestBubbleSortAllSwapped(t *testing.T) {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	Sort(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], arr[i])
		}
	}
}
