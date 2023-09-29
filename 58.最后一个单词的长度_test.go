/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 *
 * https://leetcode.cn/problems/length-of-last-word/description/
 *
 * algorithms
 * Easy (43.53%)
 * Likes:    635
 * Dislikes: 0
 * Total Accepted:    474.8K
 * Total Submissions: 1.1M
 * Testcase Example:  '"Hello World"'
 *
 * 给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。
 *
 * 单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "Hello World"
 * 输出：5
 * 解释：最后一个单词是“World”，长度为5。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "   fly me   to   the moon  "
 * 输出：4
 * 解释：最后一个单词是“moon”，长度为4。
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "luffy is still joyboy"
 * 输出：6
 * 解释：最后一个单词是长度为6的“joyboy”。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 10^4
 * s 仅有英文字母和空格 ' ' 组成
 * s 中至少存在一个单词
 *
 *
 */
package sf

// @lc code=start
func lengthOfLastWord(s string) int {
	k, a := 0, 0
	n := len(s)
	for i := n - 2; i >= 0; i-- {

		if s[i] == ' ' && s[i+1] != ' ' {
			k = i + 1
			break
		}
		if s[i+1] == ' ' {
			a += 1
		}
	}

	return n - k - a
}

// @lc code=end
