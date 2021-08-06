package l0277

/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
func solution(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {
		celebrityCandidate := 0
		for i := 0; i < n; i++ {
			if knows(celebrityCandidate, i) {
				celebrityCandidate = i
			}
		}
		for i := 0; i < n; i++ {
			if i == celebrityCandidate {
				continue
			}
			// 如果有一个人不满足关系
			if knows(celebrityCandidate, i) || !knows(i, celebrityCandidate) {
				return -1
			}
		}
		return celebrityCandidate
	}
}
