package algorithm

func myPow(x float64, n int) float64 {
	N := n
	if N < 0 {
		x = 1 / x
		N = -N
	}

	return fastPow(x, N)
}

func fastPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}

	half := fastPow(x, n/2)
	if n%2 == 1 {
		// 奇数
		return half * half * x
	} else {
		// 偶数
		return half * half
	}
}
