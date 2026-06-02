package v1

func (v Crypto) Encode(data []byte, key string) []byte {
	keyBytes, keySum := v.keygen(key)
	keyLen := uint16(len(key))
	encoded := make([]byte, len(data))

	// If len(data) = 7 and len(key) = 3
	// i: 0 1 2 3 4 5 6
	// j: 0 1 2 0 1 2 0 (Circular indexing)
	for i := range data {
		j := uint16(i) % keyLen

		// Formula: (T[i] + keySum + K[j]) % 256
		res := (uint16(data[i]) + keySum + uint16(keyBytes[j])) % 256
		encoded[i] = byte(res)
	}

	return encoded
}
