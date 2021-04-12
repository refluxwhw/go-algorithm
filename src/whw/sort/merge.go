package sort

type merge struct{}

const max = 64

func (s *merge) sortMinArr(arr []int) []int {
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

func (s *merge) sort(arr []int) (ret []int) {
	idx := len(arr) / 2
	if len(arr) > max {
		return s.merge(s.sort(arr[:idx]), s.sort(arr[idx:]))
	} else {
		return s.merge(s.sortMinArr(arr[:idx]), s.sortMinArr(arr[idx:]))
	}
	//if len(arr) > 2 {
	//	return s.merge(s.sort(arr[:idx]), s.sort(arr[idx:]))
	//} else if len(arr) == 1 {
	//	return arr
	//}
	//if arr[0] > arr[1] {
	//	arr[0], arr[1] = arr[1], arr[0]
	//}
	//return arr
}

func (s *merge) merge(arr1, arr2 []int) (ret []int) {
	ret = make([]int, 0, len(arr1)+len(arr2))

	i, j := 0, 0
	for ; i < len(arr1) && j < len(arr2); {
		if arr1[i] < arr2[j] {
			ret = append(ret, arr1[i])
			i++
		} else {
			ret = append(ret, arr2[j])
			j++
		}
	}

	if i == len(arr1) { // arr1 为空
		ret = append(ret, arr2[j:]...)
	} else {
		ret = append(ret, arr1[i:]...)
	}

	return
}

func (s *merge) Sort(arr []int) []int {
	return s.sort(arr)
}
