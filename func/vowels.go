package core

func Vowels(str []string) []string {
	for i := 0; i < len(str); i++ {
		if str[i] == "a" || str[i] == "A" {
			vowels := []string{"a", "e", "o", "i", "u", "h", "A", "E", "O", "I", "U", "H"}

			if i+1 < len(str) {
				for _, letter := range vowels {
					if letter == string(str[i+1][0]) {
						if str[i] == "a" {
							str[i] = "an"
						} else if str[i] == "A" {
							str[i] = "An"
						}
						break
					}
				}
			}
		}
	}
	return str
}
