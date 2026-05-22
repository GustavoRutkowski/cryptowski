package bruteforce

import "github.com/GustavoRutkowski/cryptowski/internal/cryptowski"

type Keysize struct {
	Min uint8
	Max uint8
}

func Ranged(data []byte, keyrange Keysize, decodeAlg cryptowski.DecodeAlg) (string, []byte, error) {
	for i := keyrange.Min; i <= keyrange.Max; i++ {
		key, decoded, err := Fixed(data, i, decodeAlg)
		
		if err != nil {
			return "", nil, err
		}

		if decoded != nil {
			return key, decoded, nil
		}
	}

	return "", nil, nil
}
