package main

import (
	"fmt"

	"github.com/pratikms/deque/deque"
)

func main() {
	d := deque.NewDeque[string](5)

	fmt.Printf("Deque front: %+v\n", d.Front())
	fmt.Printf("Deque back: %+v\n", d.Back())

	d.PushFront("rafta")
	d.PushFront("rafta")
	d.PushFront("haule")

	fmt.Printf("Deque: %+v\n", d.GetElements())

	d.PushBack("haule")
	d.PushBack("dil ko")

	fmt.Printf("Deque: %+v\n", d.GetElements())

	d.PopFront()

	fmt.Printf("Deque: %+v\n", d.GetElements())

	d.PushFront("churaya tumne")

	fmt.Printf("Deque: %+v\n", d.GetElements())

	d.PushFront("dil ko to pata bhi na chala")

	d.Front()
	d.Back()

	fmt.Printf("Deque: %+v\n", d.GetElements())
}
