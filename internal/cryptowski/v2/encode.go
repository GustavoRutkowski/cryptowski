package v2

func (v Crypto) Encode(data []byte, key string) []byte {
	padded := pkcs7{}.padding(data)
	blocks := blocksFromBytes(padded)

	k := keyToUint64(key)
	subs := keySchedule(k)

	outBlocks := make([]uint64, len(blocks))
	for i, b := range blocks {
		outBlocks[i] = feistelNetwork(b, subs)
	}

	return bytesFromBlocks(outBlocks)
}
