package sf_test

import (
	"fmt"
	"testing"
)

func TestXXX(t *testing.T) {
	fmt.Println(
		removeDuplicates([]int{1, 1, 2, 2, 3}),
	)
}
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	a, b := 2, 2
	for b < n {
		if nums[a-2] != nums[b] {
			nums[a] = nums[b]
			a++
		}
		b++
	}
	return a
}
