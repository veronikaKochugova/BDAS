package obfuscation

import (
	"strings"
)

const (
	sourceAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	targetAlpha = "QAZWSXEDCRFVTGBYHNUJMIKOLP"
	sourceDigh  = "0123456789"
	targetDigh  = "2534678901"
)

var (
	source = sourceAlpha + strings.ToLower(sourceAlpha) + sourceDigh
	target = targetAlpha + strings.ToLower(targetAlpha) + targetDigh
)

func applyMask(data, src, trg string) (result string) {
	for _, r := range data {
		i := strings.Index(src, string(r))
		if i != -1 {
			result += string(trg[i])
		} else {
			result += string(r)
		}
	}
	return result
}

func Obfuscate(s string) (result string) {
	return applyMask(s, source, target)
}

func Deobfuscate(s string) (result string) {
	return applyMask(s, target, source)
}
