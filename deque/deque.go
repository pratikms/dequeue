package deque

import "fmt"

type Deque[T any] struct {
	elements []T
	size     int
}

func (d *Deque[T]) GetLength() int {
	return len(d.elements)
}

func (d *Deque[T]) IsFull() bool {
	return d.GetLength() == d.size
}

func (d *Deque[T]) IsEmpty() bool {
	return d.GetLength() == 0
}

func (d *Deque[T]) PushBack(elem T) {
	if d.IsFull() {
		fmt.Printf("Cannot push %v from back. Queue overflow\n", elem)
		return
	}
	d.elements = append(d.elements, elem)
}

func (d *Deque[T]) PushFront(elem T) {
	if d.IsFull() {
		fmt.Printf("Cannot push %v from front. Queue overflow\n", elem)
		return
	}
	d.elements = append([]T{elem}, d.elements...)
}

func (d *Deque[T]) PopBack() T {
	var elem T
	if d.IsEmpty() {
		fmt.Println("Cannot pop form back. Queue underflow")
		return elem
	}

	idx := d.GetLength() - 1
	elem = d.elements[idx]
	d.elements = d.elements[:idx]

	return elem
}

func (d *Deque[T]) PopFront() T {
	var elem T
	if d.IsEmpty() {
		fmt.Println("Cannot pop from front. Queue underflow")
		return elem
	}

	elem = d.elements[0]
	d.elements = d.elements[1:]

	return elem
}

func (d *Deque[T]) Front() T {
	var elem T
	if d.IsEmpty() {
		return elem
	}
	return d.elements[0]
}

func (d *Deque[T]) Back() T {
	var elem T
	if d.IsEmpty() {
		return elem
	}
	return d.elements[d.GetLength()-1]
}

func (d *Deque[T]) GetElements() []T {
	return d.elements
}

func NewDeque[T any](size int) *Deque[T] {
	return &Deque[T]{
		size: size,
	}
}
