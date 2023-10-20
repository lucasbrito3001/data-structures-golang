package datastructure

type Queue struct{ Values []int }

func (q *Queue) Enqueue(element int) bool {
	q.Values = append(q.Values, element)

	return true
}

func (q *Queue) Dequeue() bool {
	if len(q.Values) == 0 {
		return false
	}

	q.Values = q.Values[1 : len(q.Values)]

	return true
}

func (q *Queue) GetTail() int {
	if len(q.Values) > 0 {
		return q.Values[0]
	} else {
		return -1
	}
}

func (q *Queue) GetHead() int {
	if len(q.Values) > 0 {
		return q.Values[len(q.Values)-1]
	} else {
		return -1
	}
}

func (q *Queue) GetQueue() []int {
	return q.Values
}
