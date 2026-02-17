package recursion

import "fmt"

func printName(i, n int) {
	if i == 0 || i > n {
		fmt.Printf("i is %v", i)
		return
	}
	fmt.Printf("Name is Priyansh with i is %v\n", i)
	i++
	printName(i, n)
}

func RecPrintName() {
	n := 5
	i := 1
	printName(i, n)
}

func PrintLinearAscending() {
	recPrintNumber(1, 5)
}

func recPrintNumber(i, n int) {
	if i > n {
		return
	}
	fmt.Printf("i is %v\n", i)
	recPrintNumber(i+1, n)
}

func PrintBackTrack() {
	n := 3
	recBackTrackingLinearAsc(n, n)
}
func recBackTrackingLinearAsc(i, n int) {
	if i < 1 {
		return
	}
	recBackTrackingLinearAsc(i-1, n)
	fmt.Printf("here is %v\n", i)
}

func PrintDescBackTrack() {
	i := 1
	n := 3
	recBackTrackingLinearDesc(i, n)
}
func recBackTrackingLinearDesc(i, n int) {
	if i > n {
		return
	}
	recBackTrackingLinearDesc(i+1, n)
	fmt.Printf("here is %v\n", i)
}

func SumofFirstN() int {
	sum := 0
	n := 3
	funRec(n, sum)
	return sum
}
func funRec(i, sum int) {
	if i < 1 {
		fmt.Printf("sum is %v\n", sum)
		return
	}
	funRec(i-1, sum+i)
	// return sum
}
