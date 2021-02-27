package pkg

import "strings"

// SQLColumnToHumpStyle sql转换成驼峰模式
func SQLColumnToHumpStyle(in string) (ret string) {
	for i := 0; i < len(in); i++ {
		if i > 0 && in[i-1] == '_' && in[i] != '_' {
			s := strings.ToUpper(string(in[i]))
			ret += s
		} else if in[i] == '_' {
			continue
		} else {
			ret += string(in[i])
		}
	}
	return
}
