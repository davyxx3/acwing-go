package search

// * 二分查找可以寻找左右两个区域的分界点

// * 寻找右分界点
func right_binary_search(l int, r int) {
	for l < r {
		mid := (l + r) >> 1
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
}

// * 寻找左分界点
func left_binary_search(l int, r int) {
	for l < r {
		// ! 注意寻找左分界点，计算mid时必须+1，否则会死循环
		mid := (l + r + 1) >> 1
		if check(mid) {
			l = mid
		} else {
			r = mid - 1
		}
	}
}

// * 判断mid是否符合区域的性质
func check(i int) bool {
	return true
}
