package core

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func array2(n, m int) [][]int {
	x := make([][]int, n)
	for i := 0; i < n; i++ {
		x[i] = make([]int, m)
	}
	return x
}

func array3(n, m, k int) [][][]int {
	x := make([][][]int, n)
	for i := 0; i < n; i++ {
		x[i] = make([][]int, m)
		for j := 0; j < m; j++ {
			x[i][j] = make([]int, k)
		}
	}
	return x
}

func init3d(n, m, k int) *[][][]int {
	x := make([][][]int, n)
	for i := 0; i < n; i++ {
		x[i] = make([][]int, m)
		for j := 0; j < m; j++ {
			x[i][j] = make([]int, k)
		}
	}
	return &x
}
