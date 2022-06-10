package sorting

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	SelectionSort(randomArray)
	fmt.Println("排序结果：",randomArray)
}
