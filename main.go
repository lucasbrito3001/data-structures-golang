package main

import (
	"fmt"

	binarysearch "github.com/lucasbrito3001/binarysearch/packages/algorithms"
	datastructure "github.com/lucasbrito3001/binarysearch/packages/data-structure"
)

func runBinarySearch(list []int, target int) {
	index := binarysearch.BinarySearch(list, target)
	fmt.Printf("\nThe index of %d is %d", target, index)
}

func runQueue(values []int) {
	queue := datastructure.Queue{Values: values}

	fmt.Printf("\nInitial queue: %d", queue.GetQueue())

	for i := 0; i < 2; i++ {
		queue.Enqueue((i + 1))
		fmt.Printf("\nEnqueued: %d", queue.GetQueue())
	}

	for i := 0; i < 3; i++ {
		result := queue.Dequeue()
		if result {
			fmt.Printf("\nDequeued: %d", queue.GetQueue())
		} else {
			fmt.Printf("\nEmpty queue, can't dequeue")
		}
	}

	for i := 0; i < 2; i++ {
		queue.Enqueue(i)
		fmt.Printf("\nEnqueued: %d", queue.GetQueue())
	}

	fmt.Printf("\nTail: %d", queue.GetTail())
	fmt.Printf("\nHead: %d", queue.GetHead())
}

func runStack(values []datastructure.Element) {
	stack := datastructure.Stack{Values: values}

	fmt.Printf("\nStack values: %+v", stack.PrintStack())
	fmt.Printf("\nPush result: %t", stack.Push(datastructure.Element{Name: "Bianca", Age: 23}))
	fmt.Printf("\nStack values: %+v", stack.PrintStack())
	element, err := stack.Peek()
	if err == -1 {
		fmt.Printf("\nPeek result: EMPTY")
	} else {
		fmt.Printf("\nPeek result: %+v", element)
	}
	fmt.Printf("\nPop result: %t", stack.Pop())
	fmt.Printf("\nStack values: %+v", stack.PrintStack())
	fmt.Printf("\nPop result: %t", stack.Pop())
	fmt.Printf("\nStack values: %+v", stack.PrintStack())
	fmt.Printf("\nPop result: %t", stack.Pop())
	fmt.Printf("\nStack values: %+v", stack.PrintStack())
	fmt.Printf("\nIsEmpty result: %t", stack.IsEmpty())
	element, err = stack.Peek()
	if err == -1 {
		fmt.Printf("\nPeek result: EMPTY")
	} else {
		fmt.Printf("\nPeek result: %+v", element)
	}
}

func runSet(values []string) {
	set := datastructure.Set{Values: values}

	fmt.Printf("\nAdd result: %t", set.Add("bianca"))
	fmt.Printf("\nAdd result: %t", set.Add("bianca"))
	fmt.Printf("\nSet values: %s", set.Get())
	fmt.Printf("\nRemove result: %t", set.Remove("lucas"))
	fmt.Printf("\nSet values: %s", set.Get())
	fmt.Printf("\nRemove result: %t", set.Remove("rafael"))
	fmt.Printf("\nSet values: %s", set.Get())
	fmt.Printf("\nHas result: %t", set.Has("bianca"))
	fmt.Printf("\nHas result: %t", set.Has("lucas"))
}

func main() {
	// binary search
	fmt.Printf("STARTING BINARY SEARCH")
	target := 4
	list := []int{1, 2, 3, 5, 4}
	runBinarySearch(list, target)

	// queue
	fmt.Printf("\n\nSTARTING QUEUE")
	initialQueue := []int{}
	runQueue(initialQueue)

	// stack
	fmt.Printf("\n\nSTARTING STACK")
	initialStack := []datastructure.Element{{Name: "Lucas", Age: 23}}
	runStack(initialStack)

	// set
	fmt.Printf("\n\nSTARTING SET")
	values := []string{"lucas"}
	runSet(values)
}
