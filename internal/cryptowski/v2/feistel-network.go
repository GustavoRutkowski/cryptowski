package v2

func round(left, right uint32, subkey48 uint64) (uint32, uint32) {
	nextLeft := right                      // L' = R
	nextRight := left ^ f(right, subkey48) // R' = L XOR F(R, K')
	return nextLeft, nextRight
}

// feistelNetwork applies the rounds using the provided subkeys in order.
func feistelNetwork(block uint64, subkeys [16]uint64) uint64 {
	// Split the 64-bit block into two 32-bit halves
	left := uint32(block >> 32)
	right := uint32(block << 32 >> 32)

	for _, k := range subkeys {
		left, right = round(left, right, k)
	}

	return (uint64(left) << 32) | uint64(right) // Merge
}
