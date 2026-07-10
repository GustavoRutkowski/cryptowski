package v2

const FEISTEL_ROUNDS uint8 = 16

func round(left, right uint32, subkey48 uint64) (uint32, uint32) {
	nextLeft := right                      // L' = R
	nextRight := left ^ f(right, subkey48) // R' = L XOR F(R, K')
	return nextLeft, nextRight
}

func feistelNetwork(block uint64, subkey48 uint64) uint64 {
	// Split the 64-bit block into two 32-bit halves
	left := uint32(block >> 32)
	right := uint32(block << 32 >> 32)

	for range FEISTEL_ROUNDS {
		left, right = round(left, right, subkey48)
	}

	return (uint64(left) << 32) | uint64(right) // Merge
}
