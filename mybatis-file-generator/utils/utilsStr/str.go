package utilsStr

import "strings"

func ToCamelStr(_str string) string {
	var text string
	//for _, p := range strings.Split(name, "_") {
	for _, p := range strings.Split(_str, "_") {
		// 字段首字母大写的同时, 是否要把其他字母转换为小写
		switch len(p) {
		case 0:
		case 1:
			text += strings.ToUpper(p[0:1])
		default:
			// 字符长度大于1时
			text += strings.ToUpper(p[0:1]) + p[1:]
		}
	}
	return text

}

func ToCamelStr2(_str string) string {
	var text string
	//for _, p := range strings.Split(name, "_") {
	for i, p := range strings.Split(_str, "_") {
		// 字段首字母大写的同时, 是否要把其他字母转换为小写
		switch len(p) {
		case 0:
		case 1:
			text += strings.ToLower(p[0:1])
		default:
			// 字符长度大于1时

			if i == 0 {
				text += strings.ToLower(p[0:1]) + p[1:]

			} else {
				text += strings.ToUpper(p[0:1]) + p[1:]
			}
		}
	}
	return text

}
