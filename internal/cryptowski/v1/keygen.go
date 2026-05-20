package v1

import (
	"math/big"

	"github.com/GustavoRutkowsik/cryptowski/internal/utils"
)

func findPrime(len uint8) *big.Int {
	if len <= 0 {
		return big.NewInt(0)
	}
	if len == 1 {
		return big.NewInt(2)
	}

	start := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(len-1)), nil)
	if new(big.Int).Mod(start, big.NewInt(2)).Cmp(big.NewInt(0)) == 0 {
		start.Add(start, big.NewInt(1))
	}

	for {
		if start.ProbablyPrime(50) {
			return start
		}
		start.Add(start, big.NewInt(2))
	}
}

func digitsOf(num uint8) uint8 {
	if num == 0 {
		return 1
	}

	var count uint8 = 0
	for num > 0 {
		num /= 10
		count++
	}
	return count
}

func (v Crypto) Keygen(key string) ([]*big.Int, *big.Int) {
	sum := big.NewInt(0)
	generatedKey := make([]*big.Int, len(key))

	for i, char := range []byte(key) {
		mappedChar := utils.ByteToBigint(char)
		prime := findPrime(digitsOf(char))
		generatedKey[i] = new(big.Int).Mul(mappedChar, prime)
		sum.Add(sum, big.NewInt(int64(char)))
	}

	return generatedKey, sum
}
