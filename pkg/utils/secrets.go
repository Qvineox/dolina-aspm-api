package utils

import "regexp"

var re = regexp.MustCompile(`[^!@#$%^&*()=+\[\]{};:'",.<>\?` + "`" + `~]+`)

// MaskSecret replaces all fields with "xxx".
// If "full" is not true, masks all fields except special characters
func MaskSecret(secret string, full bool) string {
	return re.ReplaceAllStringFunc(secret, func(match string) string {
		return "xxx"
	})
}
