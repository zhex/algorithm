package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var arr = []int{5, 6, 81, 30, 40, 55, 78, 7, 99, 33}
var sorted = []int{5, 6, 7, 30, 33, 40, 55, 78, 81, 99}

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

func TestShellSort(t *testing.T) {
	data := SortableIntArray(clone(arr))
	ShellSort(data)
	assert.EqualValues(t, data, sorted)
}
