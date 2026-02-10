package main

import (
	"fmt"
	"go_tutorial/sorting"
)

func main() {
	fmt.Println("Hello Main")
	arr := []int{5, 3, 8, 4}
	sort := sorting.BubbleSortStruct{}
	sort.Sort(arr)
	fmt.Printf("Sorted array is %v", arr)
}
