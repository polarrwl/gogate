package utils

import (
	"fmt"
	"path"
	"strings"
)

/**
 * URL匹配
 */
func MatchUrl(pattern string, target string) bool {
	index01 := strings.Index(target, "?")
	if index01 > 0 {
		target = target[0:index01]
	}
	index02 := strings.Index(target, "#")
	if index02 > 0 {
		target = target[0:index02]
	}
	index03 := strings.Index(pattern, "**")
	if index03 > 0 {
		pattern01 := pattern[0:index03]
		if strings.Index(target, pattern01) == 0 {
			return true
		}
	}
	flag, err := path.Match(pattern, target)
	if err != nil {
		fmt.Println("MatchUrl匹配错误:", err)
	}
	return flag
}
