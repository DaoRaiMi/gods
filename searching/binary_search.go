package searching

const (
	recordNotFound = -1
)
/*
二分查找的前提是目标数组必须是有序的。
 */

func BinarySearch(s []int, n int) int {
	low,high := 0, len(s)-1

	for low <= high {
		mid := (low + high)/2
		if n < s[mid] {
			high = mid-1
		}else if n > s[mid] {
			low = mid+1
		}else {
			return mid
		}
	}
	return recordNotFound
}