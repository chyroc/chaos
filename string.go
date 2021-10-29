package chaos

import (
	"strings"

	"github.com/chyroc/chaos/internal/cstring"
)

// FindLastAfterSubstr 寻找字符串 s 中，匹配最后一个 substr 后的字符串
//
// 当 substr 为空的时候，必定返回空字符串
// 当 s==1/2/3 substr=/ 返回 3
func FindLastAfterSubstr(s, substr string) string {
	return cstring.FindLastAfterSubstr(s, substr)
}

// CountPrefix 和 strings.Count 类似，不过 CountPrefix 只会统计以 substr 开头的个数
//
// 当 s=**1** substr=* 的时候，返回 2
// 当 substr 为空的时候，一定返回 0
func CountPrefix(s, substr string) int {
	return cstring.CountPrefix(s, substr)
}

// ContainAny contain any of list
func ContainAny(s string, substrList []string) bool {
	for _, substr := range substrList {
		if strings.Contains(s, substr) {
			return true
		}
	}
	return false
}
