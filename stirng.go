package chaos

import (
	"github.com/chyroc/chaos/internal/cstring"
)

// FindLastAfterSubstr 寻找字符串 s 中，匹配最后一个 substr 后的字符串
//
// 当 substr 为空的时候，必定返回空字符串
// 当 s==1/2/3 substr=/ 返回 3
func FindLastAfterSubstr(s, substr string) string {
	return cstring.FindLastAfterSubstr(s, substr)
}
