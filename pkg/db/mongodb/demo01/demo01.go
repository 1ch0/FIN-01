package demo01

import "regexp"

func hidePass(str string) string {
	reg := regexp.MustCompile(`(^mongodb://.+?:)(.+)(@.+$)`)
	return reg.ReplaceAllString(str, `${1}xxx${3}`)
}
