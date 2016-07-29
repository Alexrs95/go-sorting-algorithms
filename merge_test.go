package sorting

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	list := []int{3, 5, 7, 5, 6, 2, 1}
	ordered := []int{1, 2, 3, 5, 5, 6, 7}
	MergeSort(list)
	if !reflect.DeepEqual(list, ordered) {
		t.Error("Expected ", ordered, " got ", list)
	}
}
