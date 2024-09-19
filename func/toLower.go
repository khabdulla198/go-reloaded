package core

func ToLower(s string) string {
	arr := []rune(s)

	for i := 0; i < len(s); i++ {
		if arr[i] >= 'A' && arr[i] <= 'Z' {
			arr[i] = arr[i] + 32
		}
	}

	return string(arr)
}