package deque

import (
	"testing"
)

func TestDeque(t *testing.T) {
	deque := NewDeque[int](3)

	// Pushing elements
	deque.PushBack(1)
	deque.PushFront(2)
	deque.PushBack(3)

	// Pushing elements when deque is full
	deque.PushBack(4)
	deque.PushFront(5)

	// Checking elements
	if deque.Front() != 2 {
		t.Errorf("Front(): expected %v, got %v", 2, deque.Front())
	}

	if deque.Back() != 3 {
		t.Errorf("Back(): expected %v, got %v", 3, deque.Back())
	}

	if len(deque.GetElements()) != 3 {
		t.Errorf("GetElements(): expected %v, got %v", []int{2, 1, 3}, deque.GetElements())
	}

	// Popping elements
	elem1 := deque.PopBack()
	if elem1 != 3 {
		t.Errorf("PopBack(): expected %v, got %v", 3, elem1)
	}

	elem2 := deque.PopFront()
	if elem2 != 2 {
		t.Errorf("PopFront(): expected %v, got %v", 2, elem2)
	}

	elem3 := deque.PopBack()
	if elem3 != 1 {
		t.Errorf("PopBack(): expected %v, got %v", 1, elem3)
	}

	// Popping elements when deque is empty
	elem4 := deque.PopBack()
	if elem4 != 0 {
		t.Errorf("PopBack(): expected %v, got %v", 0, elem4)
	}

	elem5 := deque.PopFront()
	if elem5 != 0 {
		t.Errorf("PopFront(): expected %v, got %v", 0, elem5)
	}
}

func TestEmptyDeque(t *testing.T) {
	d := NewDeque[string](3)

	// Test IsEmpty()
	if d.IsEmpty() != true {
		t.Errorf("Expected deque to be empty")
	}

	// Test PopBack() on empty deque
	elem := d.PopBack()

	if elem != "" {
		t.Errorf("Expected PopBack() to return empty string")
	}

	// Test PopFront() on empty deque
	elem = d.PopFront()

	if elem != "" {
		t.Errorf("Expected PopFront() to return empty string")
	}

	// Test Front() on empty deque
	elem = d.Front()

	if elem != "" {
		t.Errorf("Expected Front() to return empty string")
	}

	// Test Back() on empty deque
	elem = d.Back()

	if elem != "" {
		t.Errorf("Expected Back() to return empty string")
	}

	// Test GetElements() on empty deque
	elements := d.GetElements()

	if len(elements) != 0 {
		t.Errorf("Expected GetElements() to return empty slice")
	}
}

func TestFullDeque(t *testing.T) {
	d := NewDeque[float64](2)

	// Test PushBack() on full deque
	d.PushBack(1.0)
	d.PushBack(2.0)
	d.PushBack(3.0)

	if len(d.GetElements()) != 2 {
		t.Errorf("Expected deque to be of length 2")
	}

	// Test PushFront() on full deque
	d.PushFront(0.5)

	if len(d.GetElements()) != 2 {
		t.Errorf("Expected deque to be of length 2")
	}
}
