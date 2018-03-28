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
// 在小的子序上使用插入算法，然后用归并整合，可以更好的提高效率
// 自顶向下的递归需要 1/2 * N * lgN  至 2 * N * lgN 次比较
// 耗时： 和 NlgN 成正比
// 优点： 可以处理百万级别甚至更大规模的数组
// 缺点： 需要额外空间，空间大小和 N 成正比
func MergeIntSort(data SortableIntArray) {
	mergeIntSort(data, 0, len(data)-1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MergeBottomIntSort - 自底向上归并排序法
// 自顶向下的递归需要 1/2 * N * lgN  至 N * lgN 次比较， 最多 6NlgN 次访问
// 当数组长度为 2 的幂时， 自上而下和自下而上归并的比较次数和访问次数相等，只是顺序不同
// 自底向上排序比较适合链表结构排序， 直接改变链接，不需要创建新的节点
func MergeBottomIntSort(data SortableIntArray) {
	l := data.Len()
	for sz := 1; sz < l; sz += sz {
		for lo := 0; lo < l; lo += sz * 2 {
			merge(data, lo, lo+sz-1, min(lo+sz*2-1, l-1))
		}
	}
}
