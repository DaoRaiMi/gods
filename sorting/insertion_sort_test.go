package sorting

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	InsertionSort(randomArray)
	if !sliceEqual(randomArray, sortedArray) {
		t.Errorf("failed to execute quickSort, want: %v, got: %v\n", sortedArray, randomArray)
	}
}
