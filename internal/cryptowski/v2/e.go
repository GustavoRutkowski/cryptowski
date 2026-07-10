package v2

var eTable = [48]uint8{
	32, 1,  2,  3,  4,  5,
	4,  5,  6,  7,  8,  9,
	8,  9,  10, 11, 12, 13,
	12, 13, 14, 15, 16, 17,
	16, 17, 18, 19, 20, 21,
	20, 21, 22, 23, 24, 25,
	24, 25, 26, 27, 28, 29,
	28, 29, 30, 32, 32, 01,
}

func e(right uint32) uint64 {
	const BLOCK_SIZE uint8 = 32
	const LEFT_PAD uint8 = 16
	var result uint64 = 0

	for _, position := range eTable {
		bit := uint64(right >> (BLOCK_SIZE - position)) & 1
		result = (result << 1) | bit
	}

	return result >> LEFT_PAD // shift to the right to remove the 16 leftmost bits, which are always 0
}