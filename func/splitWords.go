package core

func SplitWhiteSpaces(s string) []string {
	str := ""
	arr := []string{}

	if len(s) == 0 {
		return []string{}
	}

	for i := 0; i < len(s); i++ {

		if s[i] == ' ' || s[i] == '\n' || s[i] == '\t' {
			if str != "" {
				arr = append(arr, str)
				str = ""
			}
		} else {
			str = str + string(s[i])
		}

		if i == len(s)-1 {
			arr = append(arr, str)
		}
	}

	return arr
}
