package main

import (
	"go_tutorial/sorting"
)

func main() {

	arr := []int{5, 3, 8, 4}

	var a sorting.SortingImpl
	a = &sorting.BubbleSortStruct{}
	a.Sort(arr)

}
