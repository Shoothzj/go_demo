package demo_leetcode

func orangesRotting(grid [][]int) int {
	x := len(grid)
	y := len(grid[0])
	var queue []L0994Pair
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, L0994Pair{m: i, n: j})
			}
		}
	}
	result := -1
	for len(queue) > 0 {
		var newQueue []L0994Pair
		for i := 0; i < len(queue); i++ {
			l0994Pair := queue[i]
			m := l0994Pair.m
			n := l0994Pair.n
			// now travel this orange
			// up m-1, n
			if l0994Pair.m > 0 {
				if grid[m-1][n] == 1 {
					grid[m-1][n] = grid[m][n] + 1
					newQueue = append(newQueue, L0994Pair{m: m - 1, n: n})
				}
			}
			// down m+1, n
			if l0994Pair.m+1 < x {
				if grid[m+1][n] == 1 {
					grid[m+1][n] = grid[m][n] + 1
					newQueue = append(newQueue, L0994Pair{m: m + 1, n: n})
				}
			}
			// left m, n-1
			if l0994Pair.n > 0 {
				if grid[m][n-1] == 1 {
					grid[m][n-1] = grid[m][n] + 1
					newQueue = append(newQueue, L0994Pair{m: m, n: n - 1})
				}
			}
			// right m, n+1
			if l0994Pair.n+1 < y {
				if grid[m][n+1] == 1 {
					grid[m][n+1] = grid[m][n] + 1
					newQueue = append(newQueue, L0994Pair{m: m, n: n + 1})
				}
			}
		}
		queue = newQueue
	}
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if grid[i][j] == 1 {
				return -1
			} else if grid[i][j] > result {
				result = grid[i][j]
			}
		}
	}
	if result == 0 {
		return 0
	}
	return result - 2
}

type L0994Pair struct {
	m int
	n int
}
