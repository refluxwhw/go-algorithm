package sort

// 堆排序

type heap struct{}

func (s *heap) makeHeap(arr []int) {
	// 找到最后一个非叶子节点，调整，再找前一个非叶子节点，调整，一直循环到根节点
	for i := (len(arr) - 1 - 1) / 2; i >= 0; i-- {
		s.adjustHeap(arr, i, len(arr)-1)
	}
}

func (s *heap) adjustHeap(arr []int, idx int, max int) {
	newIdx := idx
	left := idx*2 + 1
	if left <= max {
		if arr[left] > arr[idx] {
			newIdx = left
		}
		right := left + 1
		if right <= max && arr[right] > arr[newIdx] {
			newIdx = right
		}
	}
	if newIdx != idx {
		arr[idx], arr[newIdx] = arr[newIdx], arr[idx]
		s.adjustHeap(arr, newIdx, max)
	}
}

func (s *heap) Sort(arr []int) []int {
	s.makeHeap(arr)
	idx := len(arr) - 1
	tmp := 0
	for idx > 0 {
		tmp = arr[0]
		arr[0] = arr[idx]
		arr[idx] = tmp
		idx--
		s.adjustHeap(arr, 0, idx)
	}

	return arr
}
