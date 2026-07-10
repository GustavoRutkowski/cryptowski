package v2

var pBox = [32]uint8{
	16,  7, 20, 21, 29, 12, 28, 17,
	1, 15, 23, 26,  5, 18, 31,  10,
	2,  8, 24, 14, 32, 27,  3,  9,
	19, 13, 30,  6, 22, 11,  4, 25,
}

func permute(right uint32) uint32 {
	const BLOCK_SIZE uint8 = 32
	var result uint64 = 0

	for _, position := range pBox {
		bit := uint64(right >> (BLOCK_SIZE - position) & 1)
		result = (result << 1) | bit
	}

	return uint32(result)
}
