package sort

// Sortable - 可排序队列接口
type Sortable interface {
	Len() int
	Less(a, b int) bool
	Swap(i, j int)
}

// SortableIntArray - Int 数组的可排序实现
type SortableIntArray []int

func (a SortableIntArray) Len() int {
	return len(a)
}
func (a SortableIntArray) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a SortableIntArray) Less(i, j int) bool {
	return a[i] < a[j]
}
