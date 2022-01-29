package sort

func merge_sort(a []int, l int, r int) {
	if l >= r {
		return
	}
	mid := (l + r) >> 1
	// * 左右分别递归进行归并排序
	merge_sort(a, l, mid)
	merge_sort(a, mid+1, r)

	tmp := make([]int, len(a))
	i, j, k := l, mid+1, 0

	// * 先进行排序，直到某一边走到末尾
	for ; i <= mid && j <= r; k++ {
		if a[i] <= a[j] {
			tmp[k] = a[i]
			i++
		} else {
			tmp[k] = a[j]
			j++
		}
	}

	// * 将剩下的数全部复制
	for ; i <= mid; i, k = i+1, k+1 {
		tmp[k] = a[i]
	}
	for ; j <= r; j, k = j+1, k+1 {
		tmp[k] = a[j]
	}

	// * 将数据从临时数组复制回原始数组
	for i, j := l, 0; i <= r; i, j = i+1, j+1 {
		a[i] = tmp[j]
	}
}
