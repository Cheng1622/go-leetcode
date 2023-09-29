/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 *
 * https://leetcode.cn/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (43.50%)
 * Likes:    2933
 * Dislikes: 0
 * Total Accepted:    1.2M
 * Total Submissions: 2.7M
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * 编写一个函数来查找字符串数组中的最长公共前缀。
 *
 * 如果不存在公共前缀，返回空字符串 ""。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：strs = ["flower","flow","flight"]
 * 输出："fl"
 *
 *
 * 示例 2：
 *
 *
 * 输入：strs = ["dog","racecar","car"]
 * 输出：""
 * 解释：输入不存在公共前缀。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= strs.length <= 200
 * 0 <= strs[i].length <= 200
 * strs[i] 仅由小写英文字母组成
 *
 *
 */
package sf

// @lc code=start
func longestCommonPrefix(strs []string) string {
	base := strs[0]
	for i := 0; i < len(base); i++ {
		for _, str := range strs[1:] {
			if i >= len(str) || str[i] != base[i] {
				return base[:i]
			}
		}
	}
	return base
}

// @lc code=end
