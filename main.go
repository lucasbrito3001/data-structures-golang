package main

import (
	"fmt"

	binarysearch "github.com/lucasbrito3001/binarysearch/packages/algorithms"
	datastructure "github.com/lucasbrito3001/binarysearch/packages/data-structure"
)

func main() {
	list := []int{1, 2, 3, 5, 4}
	target := 4

	// binary search
	index := binarysearch.BinarySearch(list, target)
	fmt.Println(index)

	// queue
	enqueue, dequeue, getTail, getHead, getQueue := datastructure.Queue([]int{})

	fmt.Printf("\nINITIAL QUEUE: %d", getQueue())

	for i := 0; i < 2; i++ {
		enqueue((i + 1))
		fmt.Printf("\nENQUEUED: %d", getQueue())
	}

	for i := 0; i < 3; i++ {
		result := dequeue()
		if result {
			fmt.Printf("\nDEQUEUED: %d", getQueue())
		} else {
			fmt.Printf("\nEMPTY QUEUE, CAN'T DEQUEUE")
		}
	}

	for i := 0; i < 2; i++ {
		enqueue(i)
		fmt.Printf("\nENQUEUED: %d", getQueue())
	}

	fmt.Printf("\nTAIL: %d", getTail())
	fmt.Printf("\nHEAD: %d", getHead())
}
