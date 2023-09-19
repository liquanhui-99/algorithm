package algorithm

// generateParenthesis 生成指定对数的括号，只生成()，不包括其它的括号
func generateParenthesis(n int) []string {
	res := []string{}
	var fn func(int, int, int, string)
	fn = func(left, right, n int, s string) {
		// terminator
		if left == n && right == n {
			res = append(res, s)
		}

		// process current logic: left, right

		// drill down
		if left < n {
			fn(left+1, right, n, s+"(")
		}
		if right < left {
			fn(left, right+1, n, s+")")
		}

		// clear current process
	}
	fn(0, 0, n, "")
	return res
}
