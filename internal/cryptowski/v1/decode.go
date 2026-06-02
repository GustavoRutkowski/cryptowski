package v1

func (v Crypto) Decode(data []byte, key string) []byte {
	keyBytes, keySum := v.keygen(key)
	keyLen := len(key)
	decoded := make([]byte, len(data))

	// If len(data) = 7 and len(key) = 3
	// i: 0 1 2 3 4 5 6
	// j: 0 1 2 0 1 2 0 (Circular indexing)
	for i := range data {
		j := i % keyLen

		// Formula: (T[i] - keySum - K[j]) % 256
		// Note: The % operator with negative numbers can yield negative results.
		// To ensure we get a positive result, we can add 256 to the result before applying the modulus again.
		res := (uint16(data[i]) - keySum - uint16(keyBytes[j])) % 256
		decoded[i] = byte(res)
	}

	return decoded
}
