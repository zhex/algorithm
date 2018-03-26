package sort

// ShellSort - 希尔排序法
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
