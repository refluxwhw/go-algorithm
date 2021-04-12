package sort

type radix struct{}

func (s *radix) findMax(arr []int) int {
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func (s *radix) Sort(arr []int) []int {
	// 16进制桶
	// 右移位数
	max := s.findMax(arr)
	for move := 0; (max >> move) > 0; move += 4 {
		buckets := make([][]int, 16)
		for i := 0; i < len(buckets); i++ {
			buckets[i] = make([]int, 0, len(arr)/16)
		}

		idx := 0
		for i := 0; i < len(arr); i++ {
			idx = (arr[i] >> move) % 16
			buckets[idx] = append(buckets[idx], arr[i])
		}

		idx = 0
		for i := 0; i < len(buckets); i++ {
			for j := 0; j < len(buckets[i]); j++ {
				arr[idx] = buckets[i][j]
				idx++
			}
		}
	}

	return arr
}
