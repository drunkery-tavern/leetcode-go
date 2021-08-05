/*给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
如果反转后整数超过 32 位的有符号整数的范围 [−231, 231 − 1] ，就返回 0。
假设环境不允许存储 64 位整数（有符号或无符号）。（-2147483648~2147483647）
示例 1：
输入：x = 123
输出：321
示例 2：
输入：x = -123
输出：-321
示例 3：
输入：x = 120
输出：21
示例 4：
输入：x = 0
输出：0
提示：
-231 <= x <= 231 - 1
Related Topics 数学
👍 2975 👎 0*/
package cn
import (
	"math"
	"strconv"
)
//leetcode submit region begin(Prohibit modification and deletion)
func reverse(x int) (rev int) {
	xStr := strconv.Itoa(x)
	revStr := ""
	if x < 0 {
		revStr = "-"
		xStr = strconv.Itoa(int(math.Abs(float64(x))))
	}
	for i := len(xStr) - 1; i >= 0; i-- {
		revStr += string(xStr[i])
	}
	rev, err := strconv.Atoi(revStr)
	if err != nil {
		println(err)
	}
	if rev < math.MinInt32 || rev > math.MaxInt32 {
		return 0
	}
	return rev

}
//leetcode submit region end(Prohibit modification and deletion)
