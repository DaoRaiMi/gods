package sorting

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	BubbleSort(randomArray)
	fmt.Println("排序结果：",randomArray)
}
