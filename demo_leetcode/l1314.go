package demo_leetcode

func matrixBlockSum(mat [][]int, k int) [][]int {
	m := len(mat)
	n := len(mat[0])
	// 加前方的数
	for _, array := range mat {
		aux := 0
		for i, val := range array {
			aux += val
			array[i] = aux
		}
	}
	// 加上方的数
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			mat[i][j] += mat[i-1][j]
		}
	}
	result := make([][]int, len(mat))
	for i := 0; i < m; i++ {
		array := make([]int, n)
		result[i] = array
		for j := 0; j < n; j++ {
			var left, top, right, down int
			right = j + k
			if right > n-1 {
				right = n - 1
			}
			down = i + k
			if down > m-1 {
				down = m - 1
			}
			array[j] = mat[down][right]
			top = i - k - 1
			left = j - k - 1

			if top >= 0 && left >= 0 {
				array[j] += mat[top][left]
			}
			if top >= 0 {
				array[j] -= mat[top][right]
			}
			if left >= 0 {
				array[j] -= mat[down][left]
			}
		}
	}
	return result
}
