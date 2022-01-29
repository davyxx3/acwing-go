package operations

// * c = a + b, a >= 0, b >= 0
func add(a []int, b []int) []int {
	// * 确保a的长度始终大于b，方便计算
	if len(a) < len(b) {
		return add(b, a)
	}
	res := make([]int, 0, len(a))
	t := 0
	for i := range a {
		t += a[i]
		if i < len(b) {
			t += b[i]
		}
		res = append(res, t%10)
		t /= 10
	}
	// * 当进位还剩有值时，把该值放到结果的最高位
	if t != 0 {
		res = append(res, t)
	}

	return res
}
