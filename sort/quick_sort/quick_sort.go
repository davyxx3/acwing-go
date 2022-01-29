package sort

func quickSort(a []int, l int, r int) {
	// * 如果只有一个元素，则必然是有序的，直接返回
	if l >= r {
		return
	}
	x, i, j := a[(l+r)>>1], l-1, r+1
	// * i左边的数都小于等于x
	// * j右边的数都大于等于x
	// * 所以当i和j相遇，或i超过j时，说明排序已经完成了
	for i < j {
		for {
			i++
			if a[i] >= x {
				break
			}
		}
		for {
			j--
			if a[j] <= x {
				break
			}
		}
		if i < j {
			t := a[i]
			a[i] = a[j]
			a[j] = t
		}
	}
	// 递归排序的边界用i和j都可以
	// ! 注意当边界用i时，分界点不能取a[l]，否则会有边界问题（可以取a[r]或a[(l+r)/2]）
	// ! 同理，边界用j时，分界点不能取a[r]
	// * 完成排序后，j指向的数是最后一个小于x的数，因此左递归排序的右边界为j
	quickSort(a, l, j)
	// * j右边的数都大于等于x，则右递归排序的左边界为j + 1
	quickSort(a, j+1, r)
}
