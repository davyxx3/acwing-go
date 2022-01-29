package operations

// * c = a * b, a >= 0, b >= 0
func mul(a []int, b int) []int {
	res := []int{}
	t := 0
	// * 当i没有走到数组末尾，或者进位还剩有值的情况下，i继续遍历
	for i := 0; i < len(a) || t != 0; i++ {
		if i < len(a) {
			t += a[i] * b
		}
		res = append(res, t%10)
		t /= 10
	}

	// * 去除高位的0
	for len(res) > 1 && res[len(res)-1] == 0 {
		res = res[:len(res)-1]
	}

	return res
}
