package sorting

type SelectionSortStruct struct{}

func (s *SelectionSortStruct) Sort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	sorted := s.mergeSort(arr)
	// copy sorted values back into the provided slice
	for i := range arr {
		arr[i] = sorted[i]
	}
}

func (s *SelectionSortStruct) mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := s.mergeSort(arr[:mid])
	right := s.mergeSort(arr[mid:])
	return s.merge(left, right)
}

func (s *SelectionSortStruct) merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
