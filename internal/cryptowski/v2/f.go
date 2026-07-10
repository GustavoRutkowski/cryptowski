package v2

func f(right32 uint32, subkey48 uint64) uint32 {
	right48 := e(right32) ^ subkey48 // E expansion (32 bits --> 42 bits)
	subs := substitute(right48)      // S-Boxes (48 bits --> 32 bits)
	return permute(subs)             // P-box (32 bits --> 32 bits)
}
