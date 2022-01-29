package operations

// * c = a - b, a >=b, a >= 0, b>= 0
func sub(a []int, b []int) []int {
	res := make([]int, 0, len(a))
	t := 0
	for i := range a {
		t = a[i] - t
		if i < len(b) {
			t -= b[i]
		}
		// * 结果有两种情况：未进位和进位
		// * 这两种情况的结果可以统一写为(t+10)%10
		res = append(res, (t+10)%10)
		if t < 0 {
			// * 若t为负，则说明进位了，t置为1
			t = 1
		} else {
			// * 若t为正，则无需进位，t为0
			t = 0
		}
	}

	// * 去除高位的0
	for len(res) > 1 && res[len(res)-1] == 0 {
		res = res[:len(res)-1]
	}

	return res
}
