package sort

func merge(data SortableIntArray, lo, mid, hi int) {
	i := lo
	j := mid + 1
	arr := append(SortableIntArray{}, data...)

	for k := lo; k <= hi; k++ {
		if i > mid {
			data[k] = arr[j]
			j++
		} else if j > hi {
			data[k] = arr[i]
			i++
		} else if arr.Less(j, i) {
			data[k] = arr[j]
			j++
		} else {
			data[k] = arr[i]
			i++
		}
	}
}

func mergeIntSort(data SortableIntArray, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	mergeIntSort(data, lo, mid)
	mergeIntSort(data, mid+1, hi)
	merge(data, lo, mid, hi)
}

// MergeIntSort - 归并排序法
// 自顶向下的递归需要 1/2 * N * lgN  至 2 * N * lgN 次比较
func MergeIntSort(data SortableIntArray) {
	mergeIntSort(data, 0, len(data)-1)
}
