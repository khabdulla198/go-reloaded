package core

func Join(strs []string, sep string) string {
	str := ""

	for i := 0; i < len(strs); i++ {
		if i == len(strs)-1 {
			str = str + strs[i]
		} else {
			str = str + strs[i] + sep
		}
	}
	return str
}
