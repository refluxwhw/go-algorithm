package sort

// 冒泡
type bubble struct{}

func (*bubble) Sort(arr []int) []int {
	l := len(arr)
	for i := l-1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
