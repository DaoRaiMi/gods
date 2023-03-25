package sorting

func QuickSort(nums []int) {
	quickRecursive(nums, 0, len(nums)-1)
}

func quickRecursive(nums []int, lft, rgt int) {
	if lft > rgt {
		return
	}
	pivot, left, right := nums[lft], lft, rgt
	for left < right {
		for nums[right] > pivot && left < right {
			right--
		}
		nums[left] = nums[right]

		for nums[left] < pivot && left < right {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = pivot

	quickRecursive(nums, lft, left-1)
	quickRecursive(nums, left+1, rgt)
}
