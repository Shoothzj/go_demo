package l0384

import "math/rand"

type Solution struct {
	list []int
}

func Constructor(nums []int) Solution {
	return Solution{list: nums}
}

// Reset Resets the array to its original configuration and return it.
func (s *Solution) Reset() []int {
	return s.list
}

// Shuffle Returns a random shuffling of the array.
func (s *Solution) Shuffle() []int {
	a := make([]int, len(s.list))
	copy(a, s.list)
	l := len(a)

	for i := 0; i < l; i++ {
		j := rand.Intn(l - i)
		a[i], a[j+i] = a[j+i], a[i]
	}

	return a
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
