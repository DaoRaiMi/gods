package sorting

import (
	"fmt"
	"strings"
)

var (
	randomArray = []int{9, 4, 6, 2, 7, 10, 1, 5, 11, 3, 8}
	sortedArray = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
)

func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}

	return true
}

func BubbleSort(a []int) {
	aLen := len(a)
	for i := 0; i < aLen-1; i++ { // 确保每个元素都和切片中的元素做对比
		for j := aLen - 1; j-1 >= i; j-- { // j-1>=i 是因为最前面的已经都是排好序了的，没必要再比较了。
			fmt.Printf("compare %d vs %d\t", a[j], a[j-1])
			if a[j] < a[j-1] { // 从小到大排序
				a[j], a[j-1] = a[j-1], a[j]
			}
			fmt.Printf("array = %v\n", a)
		}
		fmt.Println(strings.Repeat("=", 10))
	}
}
