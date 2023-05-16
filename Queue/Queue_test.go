package queue

import (
	"fmt"
	"testing"
)

func TestQueuePeek(t *testing.T) {
	queue := New()
	expected := "hello world"
	queue.Enqueue(expected)

	value := fmt.Sprint(Peek(queue))
	if value != expected {
		t.Errorf("Expected %s, got %s", expected, value)
	}
}

func TestDequeue(t *testing.T) {
	queue := New()
	expected := "hello world"
	queue.Enqueue(expected)

	if length != 1 {
		t.Errorf("Expected length of 1, got %d", length)
	}

	value := fmt.Sprint(queue.Dequeue())
	if value != expected {
		t.Errorf("Expected %s, got %s", expected, value)
	}
}

func TestMultipleDequeue(t *testing.T) {
	queue := New()
	expected := "hello world"
	queue.Enqueue("test")
	queue.Enqueue("test")
	queue.Enqueue("test")
	queue.Enqueue(expected)

	if length != 4 {
		t.Errorf("Expected length of 4, got %d", length)
	}

	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	value := fmt.Sprint(queue.Dequeue())
	if value != expected {
		t.Errorf("Expected %s, got %s", expected, value)
	}
}
