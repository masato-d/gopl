package training5_16

import "strings"

func customJoin(sep string, vals ...string) string {
	return strings.Join(vals, sep)
}