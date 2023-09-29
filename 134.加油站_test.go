/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] 加油站
 *
 * https://leetcode.cn/problems/gas-station/description/
 *
 * algorithms
 * Medium (50.10%)
 * Likes:    1404
 * Dislikes: 0
 * Total Accepted:    285.2K
 * Total Submissions: 569.7K
 * Testcase Example:  '[1,2,3,4,5]\n[3,4,5,1,2]'
 *
 * 在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
 *
 * 你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
 *
 * 给定两个整数数组 gas 和 cost ，如果你可以按顺序绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。如果存在解，则 保证 它是 唯一
 * 的。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
 * 输出: 3
 * 解释:
 * 从 3 号加油站(索引为 3 处)出发，可获得 4 升汽油。此时油箱有 = 0 + 4 = 4 升汽油
 * 开往 4 号加油站，此时油箱有 4 - 1 + 5 = 8 升汽油
 * 开往 0 号加油站，此时油箱有 8 - 2 + 1 = 7 升汽油
 * 开往 1 号加油站，此时油箱有 7 - 3 + 2 = 6 升汽油
 * 开往 2 号加油站，此时油箱有 6 - 4 + 3 = 5 升汽油
 * 开往 3 号加油站，你需要消耗 5 升汽油，正好足够你返回到 3 号加油站。
 * 因此，3 可为起始索引。
 *
 * 示例 2:
 *
 *
 * 输入: gas = [2,3,4], cost = [3,4,3]
 * 输出: -1
 * 解释:
 * 你不能从 0 号或 1 号加油站出发，因为没有足够的汽油可以让你行驶到下一个加油站。
 * 我们从 2 号加油站出发，可以获得 4 升汽油。 此时油箱有 = 0 + 4 = 4 升汽油
 * 开往 0 号加油站，此时油箱有 4 - 3 + 2 = 3 升汽油
 * 开往 1 号加油站，此时油箱有 3 - 3 + 3 = 3 升汽油
 * 你无法返回 2 号加油站，因为返程需要消耗 4 升汽油，但是你的油箱只有 3 升汽油。
 * 因此，无论怎样，你都不可能绕环路行驶一周。
 *
 *
 *
 * 提示:
 *
 *
 * gas.length == n
 * cost.length == n
 * 1 <= n <= 10^5
 * 0 <= gas[i], cost[i] <= 10^4
 *
 *
 */
package sf

// @lc code=start
//
//	func canCompleteCircuit(gas []int, cost []int) int {
//		maxs := math.MinInt64
//		length := len(gas)
//		a, k := 0, 0
//		for i := 0; i < length-1; i++ {
//			a = gas[i] - cost[i] + gas[i+1]
//			if a > maxs {
//				maxs = a
//				k = i
//			}
//		}
//		a = gas[length-1] - cost[length-1] + cost[0]
//		if a > maxs {
//			maxs = 0
//			k = 0
//		}
//
// gas[k]-cost[k]+
//
//	return
func canCompleteCircuit(gas []int, cost []int) int {
	left := 0  // 余量
	start := 0 // 初始选索引0作为起点
	totalGas, totalCost := 0, 0
	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
		left += gas[i] - cost[i] // 累加每次的剩余量
		if left < 0 {            // 去不了下一站，0到i都不是起点
			start = i + 1 // 把i+1作为起点
			left = 0      // 余量归零
		}
	}
	if totalGas < totalCost { // 总油量不够，问题无解
		return -1
	}
	return start // 总加油>=总耗油，必然有解。当遍历结束时，最新的start指向成功的起点

}

// @lc code=end
