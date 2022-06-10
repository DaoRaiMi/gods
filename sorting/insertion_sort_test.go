package sorting

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	InsertionSort(randomArray)
	fmt.Println("排序结果为：",randomArray)
}
