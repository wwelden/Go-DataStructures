package arraylist

import "fmt"

type ArrayList[T comparable] struct {
	arr []T
}

func (a *ArrayList[T]) IsEmpty() bool {
	return len(a.arr) == 0
}

func (a *ArrayList[T]) Length() int {
	return len(a.arr)
}

func (a *ArrayList[T]) Print() {
	length := len(a.arr)
	for i := 0; i < length; i++ {
		fmt.Print(a.arr[i], " ")
	}
	fmt.Println()
}

func (a *ArrayList[T]) Append(value T) {
	a.arr = append(a.arr, value)
}

func (a *ArrayList[T]) Remove(index int) T {
	res := a.arr[index]
	a.arr = append(a.arr[:index], a.arr[index+1:]...)
	return res
}

func (a *ArrayList[T]) Get(index int) T {
	return a.arr[index]
}

func (a *ArrayList[T]) Insert(index int, value T) {
	a.arr = append(a.arr[:index], append([]T{value}, a.arr[index:]...)...)
}


