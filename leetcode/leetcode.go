package leetcode

// nums := []int{2, 5, 4, 3}
// 	res := longestBalanced(nums)
// 	fmt.Printf("res is %v", res)

func LongestBalanced(nums []int) int {
	var maxLen int
	maxLen = 0
	n := len(nums)
	for i := 0; i < n; i++ {
		evenMap := make(map[int]struct{}, 0)
		oddMap := make(map[int]struct{}, 0)
		// fmt.Printf("nums[%v] is %v\n", i, nums[i])

		for j := i; j < n; j++ {
			num := nums[j]

			if num%2 == 0 {
				evenMap[num] = struct{}{}
			} else {
				oddMap[num] = struct{}{}
			}
			if len(evenMap) == len(oddMap) {
				length := j - i + 1
				if length > maxLen {
					maxLen = length
				}
			}
		}
	}

	return maxLen
}
