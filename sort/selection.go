package sort

// Tips:
// 1 + 2 + 3 + ... + N-1 ==  N * (N - 1) / 2

// SelectionSort - 选择排序法,
// 排序需要大约 N(N-1)/2 次比较和 N 次交换,
// 优点： 容易实现, 移动次数最少
// 缺点： 无法为下一次扫描提供帮助信息， 有序数组或者主键全部相等的数组和一个随时数组排序时间一样
func SelectionSort(data Sortable) {
	l := data.Len()

	for i := 0; i < l; i++ {
		min := i
		for j := i + 1; j < l; j++ {
			if data.Less(j, min) {
				min = j
			}
		}
		data.Swap(i, min)
	}
}
