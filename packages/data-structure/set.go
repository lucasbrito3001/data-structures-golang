package datastructure

import (
	"slices"
)

type Set struct{ Values []string }

func (s *Set) Add(element string) bool {
	for _, value := range s.Values {
		if element == value {
			return false
		}
	}

	s.Values = append(s.Values, element)

	return true
}

func (s *Set) Has(element string) bool {
	for _, value := range s.Values {
		if element == value {
			return true
		}
	}

	return false
}

func (s *Set) Remove(element string) bool {
	index := slices.Index(s.Values, element)

	if index == -1 {
		return false
	}

	s.Values = append(s.Values[:index], s.Values[index+1:]...)

	return true
}

func (s *Set) Get() []string {
	return s.Values
}
