package utils

func CheckHeader(header, expected []byte) bool {
	for i := 0; i < len(expected); i++ {
		if header[i] != expected[i] {
			return false
		}
	}
	return true
}
