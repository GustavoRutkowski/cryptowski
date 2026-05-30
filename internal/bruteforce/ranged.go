package bruteforce

import (
	"errors"
	"github.com/GustavoRutkowski/cryptowski/internal/cryptowski"
	"github.com/GustavoRutkowski/cryptowski/internal/utils"
)

func Ranged(data []byte, keyrange utils.Interval[uint8], decodeAlg cryptowski.DecodeAlg) (string, []byte, error) {
	for i := keyrange.Min; i <= keyrange.Max; i++ {
		key, decoded, err := Fixed(data, i, decodeAlg)

		if errors.Is(err, ErrKeyNotFound) {
			continue
		}

		if err != nil {
			return "", nil, err
		}

		if decoded != nil {
			return key, decoded, nil
		}
	}

	return "", nil, ErrKeyNotFound
}
