package v2

func reverseSubkeys(subs [16]uint64) [16]uint64 {
	res := [16]uint64{}
	for i := range 16 {
		res[i] = subs[16 - 1 - i]
	}
	return res
}

func (v Crypto) Decode(data []byte, key string) []byte {
	blocks := blocksFromBytes(data)
	k := keyToUint64(key)
	subs := keySchedule(k)
	
	outBlocks := make([]uint64, len(blocks))
	for i, b := range blocks {
		// For decryption the subkeys are applied in reverse order
		outBlocks[i] = feistelNetwork(b, reverseSubkeys(subs))
	}

	out := bytesFromBlocks(outBlocks)
	return pkcs7{}.unpadding(out)
}
