package operations

func div(a []int, b int) ([]int, int) {
	res := make([]int, 0, len(a))
	r := 0
	// ! 注意此处是逆序遍历，因为除法是从高位开始做的
	for i := len(a) - 1; i >= 0; i-- {
		// * 进位乘10之后加上当前的数
		r = r*10 + a[i]
		// * 进行除法运算
		res = append(res, r/b)
		// * 计算余数
		r %= b
	}
	// * 逆转结果数组（使其满足低位在前的顺序）
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		t := res[i]
		res[i] = res[j]
		res[j] = t
	}

	// * 去除高位的0
	for len(res) > 0 && res[len(res)-1] == 0 {
		res = res[:len(res)-1]
	}
	return res, r
}
