package sorting

func SelectionSort(a []int) {
	aLen := len(a)
	for i:=0;i<=aLen-1;i++{
		minIdx := i
		for j := i+1;j<=aLen-1;j++ {
			if a[minIdx] > a[j] {
				minIdx = j
			}
		}

		if a[i] > a[minIdx] {
			a[i],a[minIdx] = a[minIdx],a[i]
		}
	}
}
