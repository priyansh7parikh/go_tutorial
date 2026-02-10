package sorting

type InsertionSortStruct struct {
}

// 5,3,8,4
func (i *InsertionSortStruct) Sort(arr []int) {
	n := len(arr)

	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}
}
