package v2

import "encoding/binary"

type Crypto struct{}

const BLOCK_SIZE_BYTES = 8

func blocksFromBytes(data []byte) []uint64 {
	blocks := len(data) / BLOCK_SIZE_BYTES
	result := make([]uint64, blocks)

	for i := range blocks {
		start := i * BLOCK_SIZE_BYTES
		end := (i + 1) * BLOCK_SIZE_BYTES
		result[i] = binary.BigEndian.Uint64(data[start:end])
	}
	
	return result
}

func bytesFromBlocks(blocks []uint64) []byte {
	result := make([]byte, len(blocks) * BLOCK_SIZE_BYTES)

	for i, block := range blocks {
		start := i * BLOCK_SIZE_BYTES
		end := (i + 1) * BLOCK_SIZE_BYTES
		binary.BigEndian.PutUint64(result[start:end], block)
	}

	return result
}
