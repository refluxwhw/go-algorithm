package sort

// 插入
type insertion struct{}

func (s *insertion) Sort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		j := i
		for ; j > 0 && arr[j-1] > arr[i]; j-- {
			arr[j] = arr[j-1]
		}
		if j != i {
			tv := arr[i]
			copy(arr[j+1:i+1], arr[j:])
			arr[j] = tv
		}
	}
	return arr
}

type insertion1 struct{}

func (s *insertion1) Sort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		j := i
		tv := arr[j]
		for ; j > 0 && arr[j-1] > tv; j-- {
			arr[j] = arr[j-1]
		}
		if j != i {
			arr[j] = tv
		}
	}
	return arr
}
