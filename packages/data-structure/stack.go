package datastructure

type Element struct {
	Name string
	Age  int8
}

type Stack struct{ Values []Element }

func (s *Stack) Push(element Element) bool {
	s.Values = append(s.Values, element)
	return true
}

func (s *Stack) Pop() bool {
	if len(s.Values) == 0 {
		return false
	}

	s.Values = s.Values[1:len(s.Values)]

	return true
}

func (s *Stack) Peek() (Element, int) {
	if s.IsEmpty() {
		return Element{}, -1
	}

	return s.Values[len(s.Values)-1], 0
}

func (s *Stack) IsEmpty() bool {
	return len(s.Values) == 0
}

func (s *Stack) PrintStack() []Element {
	return s.Values
}
