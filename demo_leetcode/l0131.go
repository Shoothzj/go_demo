package demo_leetcode

func partition(s string) [][]string {
	var result = make([][]string, 0)
	current := make([]string, 0)
	dfsPartition(&result, current, s, "", 0)
	return result
}

// dfsPartition
// result dfs final result
// current dfs current
// src source string
// tail tail string
// index 当前处理到哪个Index
func dfsPartition(resultPointer *[][]string, current []string, src string, tail string, index int) {
	if index >= len(src) {
		if tail == "" {
			tmp := make([]string, len(current))
			copy(tmp, current)
			*resultPointer = append(*resultPointer, tmp)
		}
		return
	}
	newTail := tail + string(src[index])
	if canTailCommit(newTail) {
		// tail commit
		dfsPartition(resultPointer, append(current, newTail), src, "", index+1)
	}
	if index < len(src) {
		dfsPartition(resultPointer, current, src, newTail, index+1)
	}
}

func canTailCommit(src string) bool {
	if src == "" {
		return false
	}
	l := len(src)
	for i := 0; i < l/2; i++ {
		if src[i] != src[l-i-1] {
			return false
		}
	}
	return true
}
