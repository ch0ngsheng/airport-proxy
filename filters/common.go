package filters

import "strings"

// Filter 不同格式的filter需要实现的接口
type Filter interface {
	Do([]string) (error, []byte)
}

func containsKeywords(str string, keywords []string) bool {
	for _, word := range keywords {
		if strings.Contains(str, word) {
			return true
		}
	}
	return false
}
