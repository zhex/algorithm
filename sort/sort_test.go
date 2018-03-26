package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var arr = []int{5, 6, 81, 30, 40, 55}
var sorted = []int{5, 6, 30, 40, 55, 81}

func clone(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	return b
}

func TestSelectionSort(t *testing.T) {
	data := SortableIntArray(clone(arr))
	SelectionSort(data)
	assert.EqualValues(t, data, sorted)
}

func TestInsertionSort(t *testing.T) {
	data := SortableIntArray(clone(arr))
	InsertionSort(data)
	assert.EqualValues(t, data, sorted)
}
