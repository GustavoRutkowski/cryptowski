package v1

import (
	"math/big"
	"github.com/GustavoRutkowsik/cryptowski/internal/utils"
)

func (v Crypto) Encode(data []byte, key string) []byte {
	generatedKey, keySum := v.Keygen(key)
	keyLen := uint16(len(key))
	encoded := make([]byte, len(data))

	// If len(data) = 7 and len(key) = 3
	// i: 0 1 2 3 4 5 6
	// j: 0 1 2 0 1 2 0 (Circular indexing)
	for i := range data {
		j := uint16(i) % keyLen
		dataByte := utils.ByteToBigint(data[i])
		dkeyByte := generatedKey[j]

		// Formula: (T[i] + keySum + K[j]) % 256
		sum := new(big.Int).Add(dataByte, keySum)
		sum.Add(sum, dkeyByte)
		mod := new(big.Int).Mod(sum, big.NewInt(256)).Uint64()
		encoded[i] = byte(mod)
	}

	return encoded
}
