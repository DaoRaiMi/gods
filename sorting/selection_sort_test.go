package sorting

import (
	"testing"
)

func TestSelectionSort(t *testing.T) {
	SelectionSort(randomArray)
	if !sliceEqual(randomArray, sortedArray) {
		t.Errorf("failed to execute quickSort, want: %v, got: %v\n", sortedArray, randomArray)
	}
}
