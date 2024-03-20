package proc_string

import (
	"log"
	"strings"
	"tools/global"
)

func init() {
	global.Register(Upper, "string_upper")
	global.Register(Lower, "string_lower")
}

func Upper(in string) string {
	return strings.ToUpper(in)
}

func Lower(in string) string {
	return strings.ToLower(in)
}

func Length(in string) string {
	return
}
