package sorting

func InsertionSort(arr []int) {
	arrLen := len(arr)
	if arrLen <= 1 {
		return
	}

	for i := 1; i <arrLen;i++ { // 从第二个元素开始，因为默认第一个元素是有序的。
		tmp := arr[i] // 待排序元素
		var j int
		for j = i-1;j>=0;j-- { // 从已排序的后面向前面遍历
			if arr[j] > tmp { // 比较已经排序的元素与待排序元素的大小，然后再交换位置
				arr[j+1] = arr[j]
			}else {
				break
			}
		}

		// 此时位置找到了，再把待排序的元素放在这里。
		arr[j+1] = tmp // for每次循环后都会减1,所以最后一次的时候，不满足条件了，也减了1,所以要加上。
	}
}

