package sort

// InsertionSort - 插入排序
// 参考整理扑克牌时的动作， 把每张牌插入适当的位置，
// 排序时间取决于初始顺序，上来越有序，排序耗时越少
// 最优情况需要 N - 1 次比较和 0 次交换
// 最差情况 需要 ～ N * N / 2 比较和 ～ N * N / 2 交换
// 平均需要 ～ N * N / 4 比较和 N * N / 4 交换
func InsertionSort(data Sortable) {
	for i := 1; i < data.Len()-1; i++ {
		for j := i + 1; j > 0; j-- {
			if data.Less(j, j-1) {
				data.Swap(j, j-1)
			}
		}
	}
}
