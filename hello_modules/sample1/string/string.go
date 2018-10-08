package string

import (
	"strings"
)
 
func JoinWithSlash(a []string) string {
	return strings.Join(a, "/")
}
