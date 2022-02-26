package platforms

import "strings"

func fromClash(agent string) bool {
	return strings.Contains(strings.ToLower(agent), "clash")
}
