package stack

import "testing"

func TestStackPeek(t *testing.T) {
	Push("hello world")
	expected := "hello world"
	value := Peek()
	if value != expected {
		t.Errorf("Expected %s, got %s", expected, value)
	}
}

func TestStackPush(t *testing.T) {
	Push("hello world")
	Push("hello world 2")
	if length != 2 {
		t.Errorf("Expected length of 2, got %d", length)
	}
}

func TestStackPop(t *testing.T) {
	Push("hello world")
	Push("hello world 2")
	Push("hello world 3")
	Pop()
	if length != 2 {
		t.Errorf("Expected length of 2, got %d", length)
	}
}

func TestStackPopValue(t *testing.T) {
	Push("hello world")
	Push("hello world 2")
	Push("hello world 3")
	expected := "hello world 3"
	value := Pop()
	if value != expected {
		t.Errorf("Expected %s, got %s", expected, value)
	}
}

func TestMultipleStackPop(t *testing.T) {
	Push("hello world")
	Push("hello world 2")
	Push("hello world 3")
	Pop()
	Pop()
	if length != 1 {
		t.Errorf("Expected length of 1, got %d", length)
	}
}
