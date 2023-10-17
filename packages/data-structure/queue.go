package datastructure

func Queue(initialValue []int) (enqueue func(element int) bool, dequeue func() bool, getTail func() int, getHead func() int, getQueue func() []int) {
	queue := initialValue

	enqueue = func(element int) bool {
		queue = append(queue, element)

		return true
	}

	dequeue = func() bool {
		if len(queue) == 0 {
			return false
		}

		queue = queue[0 : len(queue)-1]

		return true
	}

	getTail = func() int {
		if len(queue) > 0 {
			return queue[0]
		} else {
			return -1
		}
	}

	getHead = func() int {
		if len(queue) > 0 {
			return queue[len(queue)-1]
		} else {
			return -1
		}
	}

	getQueue = func() []int {
		return queue
	}

	return
}
