package stack

import "fmt"

type Stack[T comparable] struct {
	s []T
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.s) == 0
}
func (s *Stack[T]) Length() int {
	return len(s.s)
}
func (s *Stack[T]) Print() {
	length := len(s.s)
	for i := 0; i < length; i++ {
		fmt.Print(s.s[i], " ")
	}
	fmt.Println()
}
func (s *Stack[T]) Push(value T) {
	s.s = append(s.s, value)
}
func (s *Stack[T]) Pop() T {
	length := len(s.s)
	res := s.s[length-1]
	s.s = s.s[:length-1]
	return res
}
func (s *Stack[T]) Top() T {
	length := len(s.s)
	return s.s[length-1]
}
