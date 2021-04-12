package sort

type bucket struct{}

type _bucket struct {
	arr []int
}

func (b *_bucket) insert(v int) {
	b.arr = append(b.arr, v)
	i := 0
	for ; i < len(b.arr)-1 && b.arr[i] < v; i++ {
	}
	if i != len(b.arr)-1 {
		copy(b.arr[i+1:], b.arr[i:])
		b.arr[i] = v
	}
}

func findRange(arr []int) (min, max int) {
	min, max = arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		} else if arr[i] < min {
			min = arr[i]
		}
	}
	return
}

func (s *bucket) Sort(arr []int) []int {
	const bucketNum = 64
	if len(arr) < bucketNum {
		ns := &shell{}
		return ns.Sort(arr)
	}

	min, max := findRange(arr)
	step := (max - min + bucketNum - 1) / bucketNum
	buckets := make([]_bucket, bucketNum)
	for i := 0; i < len(arr); i++ {
		idx := (arr[i] - min) / step
		if buckets[idx].arr == nil {
			buckets[idx].arr = make([]int, 0, len(arr)/bucketNum)
		}
		buckets[idx].insert(arr[i])
	}

	for idx, i := 0, 0; i < len(buckets); i++ {
		for j := 0; j < len(buckets[i].arr); j++ {
			arr[idx] = buckets[i].arr[j]
			idx++
		}
	}

	return arr
}

type bucket1 struct{}

func (s* bucket1) sort(arr []int, n int) []int  {
	if len(arr) < n * 2 {
		ns := &shell{}
		return ns.Sort(arr)
	}

	min, max := findRange(arr)
	step := (max - min) / n + 1
	buckets := make([][]int, n)
	bucketsNum := 0
	for i := 0; i < len(arr); i++ {
		idx := (arr[i] - min) / step
		if buckets[idx] == nil {
			buckets[idx] = make([]int, 0, len(arr)/n)
			bucketsNum++
		}
		buckets[idx] = append(buckets[idx], arr[i])
	}

	if bucketsNum == 1 {
		n += 2
	}

	for idx, i := 0, 0; i < len(buckets); i++ {
		a := s.sort(buckets[i], n)
		for j := 0; j < len(a); j++ {
			arr[idx] = a[j]
			idx++
		}
	}

	return arr
}

func (s *bucket1) Sort(arr []int) []int {
	return s.sort(arr, 64)
}
