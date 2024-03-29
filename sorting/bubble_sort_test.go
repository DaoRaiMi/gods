package sorting

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	BubbleSort(randomArray)
	if !sliceEqual(randomArray, sortedArray) {
		t.Errorf("failed to execute quickSort, want: %v, got: %v\n", sortedArray, randomArray)
	}
}
