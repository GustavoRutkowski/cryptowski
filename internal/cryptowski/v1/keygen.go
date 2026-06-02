package v1

func (v Crypto) keygen(key string) ([]byte, uint16) {
	var sum uint16 = 0
	bytes := []byte(key)

	for char := range bytes {
		sum += uint16(char)
	}

	return bytes, sum
}
