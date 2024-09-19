package core

func ToCap(s string) string{
	arr := []rune(s)

	for i , char := range arr{
		if i == 0 && (char >= 'a' && char <= 'z'){
			arr[i] = arr[i] - 32
		}else if i > 0 && (char >= 'A' && char <= 'Z'){
			arr[i] = arr[i] + 32
		}
	}
	return string(arr)
}