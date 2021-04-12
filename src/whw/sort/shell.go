package sort

// 希尔排序

type shell struct{}

func (s *shell) Sort(arr []int) []int {
	l := len(arr)
	for gap := l / 2; gap > 0; gap /= 2 {
		// 多个分组交替执行
		for i := gap; i < l; i++ {
			for j := i; j >= gap && arr[j-gap] > arr[j]; j -= gap {
				arr[j], arr[j-gap] = arr[j-gap], arr[j]
			}
		}
	}
	return arr
}

type shell1 struct{}

func (s *shell1) Sort(arr []int) []int {
	l := len(arr)
	for gap := l / 2; gap >= 1; gap /= 2 { // gap表示分时的下标间隔，即可以分为 gap 组
		for n := 0; n < gap; n++ { // 遍历分组，对每一组进行排序
			// 冒泡
			//for i := l - n - 1; i >= gap; i -= gap { // 从后往前，第 n 组的最后一个元素下标为 l-n-1
			//	for j := i; j < l && arr[j-gap] > arr[j]; j += gap {
			//		arr[j], arr[j-gap] = arr[j-gap], arr[j]
			//	}
			//}

			// 插入
			for i := n + gap; i < l; i += gap { // 第 n 组的第一个元素下标为 n，从第二个 n+gap 开始向前插入
				j := i
				tv := arr[j]
				for ; j >= gap && tv < arr[j-gap]; j -= gap {
					arr[j] = arr[j-gap]
				}
				if j != i {
					arr[j] = tv
				}
			}
		}
	}

	return arr
}
