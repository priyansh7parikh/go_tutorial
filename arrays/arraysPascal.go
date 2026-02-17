package arrays

import (
	"fmt"
	"sort"
)

func generate(numRows int) [][]int {
	var finalRes [][]int
	for i := 1; i < numRows; i++ {
		res := GetRowPascal2(i)
		finalRes = append(finalRes, res)
	}
	return finalRes
}

func FindElementPascal(numRows, numCols int) int {
	res := 1
	n := numRows - 1
	r := numCols - 1

	for i := 0; i < r; i++ {
		res = res * (n - 1)
		res = res / (i + 1)
	}
	fmt.Printf("res is %v", res)
	return res
}

func FindCompleteRow(numRows, numCols int) []int {
	result := make([]int, numRows)
	// default 1
	ans := 1
	// as we have index starting at 1
	n := numRows - 1
	// as we have index starting at 1 (if 0th index then it will be -2)(to optimize solution)
	r := numCols - 1

	for i := 1; i < r; i++ {
		ans = ans * (n - i)
		ans = ans / (i)
		result = append(result, ans)
	}
	return result
}

func GetRowPascal(rowIndex int) []int {
	result := make([]int, rowIndex)
	// default 1
	ans := 1
	// as we have index starting at 1
	n := rowIndex
	// as we have index starting at 1 (if 0th index then it will be -2)(to optimize solution)
	// r := numCols - 1
	result = append(result, ans)
	for i := 1; i < n; i++ {
		ans = ans * (n - i)
		ans = ans / (i)
		result = append(result, ans)
	}
	return result
}
func GetRowPascal2(rowIndex int) []int {
	result := make([]int, rowIndex)
	// default 1
	ans := 1
	// as we have index starting at 1
	n := rowIndex - 1
	// as we have index starting at 1 (if 0th index then it will be -2)(to optimize solution)
	// r := numCols - 1
	result = append(result, ans)
	for i := 1; i < n; i++ {
		ans = ans * (n - i)
		ans = ans / (i)
		result = append(result, ans)
	}
	return result
}

// [10,20,30,10]
func ToggleLightBulbs(bulb []int) []int {
	var turnedOnBulb []int
	status := make(map[int]bool, 0)
	// true if switch on , switch off --> false
	for i := 0; i < len(bulb); i++ {

		status[bulb[i]] = !status[bulb[i]]

	}
	seen := make(map[int]bool)
	for i := 0; i < len(bulb); i++ {
		if status[bulb[i]] && !seen[bulb[i]] {
			turnedOnBulb = append(turnedOnBulb, bulb[i])
			seen[bulb[i]] = true
		}
	}
	sort.Ints(turnedOnBulb)
	return turnedOnBulb
}

// func FirstUnqiueFreq(nums []int) int {

// }

func setZeroes(matrix [][]int) {
}

func Pattern() {
	// matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Printf("%v", matrix[i][j])
		}
		fmt.Printf("\n")
	}
	//1
	//12
	//123
	//1234
	//12345
}

func NextPermutation(nums []int) {
	ind := -1
	n := len(nums)
	for i := n - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			ind = i
			break
		}
	}
	if ind == -1 {

	}
}
