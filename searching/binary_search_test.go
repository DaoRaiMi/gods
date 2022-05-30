package searching

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	s := []int{0,1,16,24,35,47,59,62,73,88,99}
	fmt.Println(BinarySearch(s,62))
}
