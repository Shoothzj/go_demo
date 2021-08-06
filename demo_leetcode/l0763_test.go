package demo_leetcode

import (
	"fmt"
	"testing"
)

func TestPartitionLabels(t *testing.T) {
	labels := partitionLabels("ababcbacadefegdehijhklij")
	fmt.Println(labels)
}
