package v1

import (
	"math/big"
	"github.com/GustavoRutkowski/cryptowski/internal/utils"
)

func (v Crypto) Decode(data []byte, key string) []byte {
	generatedKey, keySum := v.Keygen(key)
	keyLen := uint16(len(key))
	decoded := make([]byte, len(data))

	// If len(data) = 7 and len(key) = 3
	// i: 0 1 2 3 4 5 6
	// j: 0 1 2 0 1 2 0 (Circular indexing)
	for i := range data {
		j := uint16(i) % keyLen
		dataByte := utils.ByteToBigint(data[i])
		dkeyByte := generatedKey[j]

		// Formula: (T[i] - keySum - K[j]) % 256
		// Note: The % operator with negative numbers can yield negative results.
		// To ensure we get a positive result, we can add 256 to the result before applying the modulus again.
		rest := new(big.Int).Sub(dataByte, keySum)
		rest.Sub(rest, dkeyByte)
		mod := new(big.Int).Mod(rest, big.NewInt(256)).Uint64()
		decoded[i] = byte(mod)
	}

	return decoded
}
