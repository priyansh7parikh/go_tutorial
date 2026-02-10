package sorting

type QuickSortStruct struct{}

func (q *QuickSortStruct) Sort(arr []int) {
	if len(arr) == 0 {
		return
	}
	q.quickSort(arr, 0, len(arr)-1)
}

func (q *QuickSortStruct) quickSort(arr []int, low, high int) {
	if low < high {
		p := q.partition(arr, low, high)
		q.quickSort(arr, low, p-1)
		q.quickSort(arr, p+1, high)
	}
}

func (q *QuickSortStruct) partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]

	return i + 1
}
