package sort

type selection struct{}

func (*selection) Sort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		idx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[idx] {
				idx = j
			}
		}
		if idx != i {
			arr[i], arr[idx] = arr[idx], arr[i]
		}
	}
	return arr
}
