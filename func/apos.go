package core

import "strings"

func Apostrophe(stri []string) string {
	//seperate apostrophe
	for i, char := range stri {
		if strings.HasPrefix(char, "'") {
			stri[i] = "' " + strings.TrimPrefix(char, "'")
		}
		if strings.HasSuffix(char, "'") {
			stri[i] = strings.TrimSuffix(char, "'") + " '"
		}
	}

	strA := strings.Join(stri, " ")
	str := strings.Fields(strA)

	c := false
	for i := 0; i < len(str); i++ {
		if str[i] == "'" && !c {
			c = true
			if i+1 < len(str) {
				str[i+1] = "'" + str[i+1]
				// str = append(str[:i], str[i+1:]...)
				// i--
				str[i] = ""
			}

		} else if str[i] == "'" && c {
			if i > 0 {
				str[i-1] = str[i-1] + "'"
				// str = append(str[:i], str[i+1:]...)
				// i--
			} else {
				str[i+1] = "'" + str[i+1]
				// str = append(str[:i], str[i+1:]...)
				// i--
			}
			str[i] = ""
			c = false
		}
	}
	strOut := strings.Join(str, " ")
	strOut = strings.Replace(strOut, " '", "'", -1)
	strOut = strings.Replace(strOut, "' ", "'", -1)

	return strOut

}
