/*
 * @lc app=leetcode.cn id=135 lang=golang
 *
 * [135] 分发糖果
 *
 * https://leetcode.cn/problems/candy/description/
 *
 * algorithms
 * Hard (50.02%)
 * Likes:    1344
 * Dislikes: 0
 * Total Accepted:    250.6K
 * Total Submissions: 501.3K
 * Testcase Example:  '[1,0,2]'
 *
 * n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。
 *
 * 你需要按照以下要求，给这些孩子分发糖果：
 *
 *
 * 每个孩子至少分配到 1 个糖果。
 * 相邻两个孩子评分更高的孩子会获得更多的糖果。
 *
 *
 * 请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：ratings = [1,0,2]
 * 输出：5
 * 解释：你可以分别给第一个、第二个、第三个孩子分发 2、1、2 颗糖果。
 *
 *
 * 示例 2：
 *
 *
 * 输入：ratings = [1,2,2]
 * 输出：4
 * 解释：你可以分别给第一个、第二个、第三个孩子分发 1、2、1 颗糖果。
 * ⁠    第三个孩子只得到 1 颗糖果，这满足题面中的两个条件。
 *
 *
 *
 * 提示：
 *
 *
 * n == ratings.length
 * 1 <= n <= 2 * 10^4
 * 0 <= ratings[i] <= 2 * 10^4
 *
 *
 */
package sf

import "testing"

// @lc code=start
// func candy(ratings []int) int {
// 	length := len(ratings)
// 	k := make([]int, length)

// 	for i := 0; i < length-1; i++ {
// 		if ratings[i] < ratings[i+1] {
// 			k[i+1] = k[i] + 1
// 		}
// 		// if ratings[i] == ratings[i+1] {
// 		// 	if k[i] == k[i+1] {
// 		// 		k[i] += 1
// 		// 	}
// 		// }
// 		if ratings[i] > ratings[i+1] && k[i] <= k[i+1] {
// 			k[i] += 1
// 		}
// 	}
// 	h := 0
// 	for _, v := range k {
// 		v += 1
// 		h += v
// 	}
// 	return h
// }

func candy(ratings []int) (ans int) {
	n := len(ratings)
	left := make([]int, n)
	for i, r := range ratings {
		if i > 0 && r > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	right := 0
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		ans += maxcandy(left[i], right)
	}
	return
}

func maxcandy(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestXXX(t *testing.T) {
	candy([]int{1, 3, 2, 2, 1})
}

// @lc code=end
