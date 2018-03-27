package sort

// ShellSort - 希尔排序法
// 把序列分成若干子序列， 对子序列应用插入排序
// 减小增量，再排序
// 最后对整个序列再做一次插入排序
// 希尔排序更高效的原因是它权衡了子数组的规则和有序性
func ShellSort(data Sortable) {
	l := data.Len()
	h := 1

	for h < l/3 {
		h = h*3 + 1
	}

	for h >= 1 {
		for i := h; i < l; i++ {
			for j := i; j >= h; j -= h {
				if data.Less(j, j-h) {
					data.Swap(j, j-h)
				}
			}
		}
		h = h / 3
	}
}
