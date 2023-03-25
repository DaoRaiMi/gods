package sorting

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	QuickSort(randomArray)
	if !sliceEqual(randomArray, sortedArray) {
		t.Errorf("failed to execute quickSort, want: %v, got: %v\n", sortedArray, randomArray)
	}
}
