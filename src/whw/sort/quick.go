package sort

// 快排

type quick struct {
}

func (s *quick) quickSort(arr []int, left, right int) {
	L := left
	R := right
	V := arr[L]
	for L < R {
		for ; L < R && arr[R] >= V; R-- {
		}
		if L < R {
			arr[L] = arr[R]
			L++
		}

		for ; L < R && arr[L] < V; L++ {
		}
		if L < R {
			arr[R] = arr[L]
			R--
		}
	}
	arr[R] = V

	if left < R-1 {
		s.quickSort(arr, left, R-1)
	}
	if right > R+1 {
		s.quickSort(arr, R+1, right)
	}
}

func (s *quick) Sort(arr []int) []int {
	s.quickSort(arr, 0, len(arr)-1)
	return arr
}
