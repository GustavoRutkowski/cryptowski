package v2

import (
	"encoding/binary"
	"math/bits"
)

func keyToUint64(key string) uint64 {
	var keyBytes [8]byte
	copy(keyBytes[:], []byte(key)) // truncates or pads with zeros
	return binary.BigEndian.Uint64(keyBytes[:])
}

// keySchedule produces 16 subkeys of 48 bits each from a 64-bit key.
// This is a simple schedule: rotate the base key and mask to 48 bits per round.
func keySchedule(key uint64) [16]uint64 {
	const MASK_48 = (uint64(1) << 48) - 1
	const ROUNDS = 16

	subs := [16]uint64{}
	for i := range ROUNDS {
		rot := bits.RotateLeft64(key, i + 1)
		subs[i] = rot & MASK_48
	}
	return subs
}
