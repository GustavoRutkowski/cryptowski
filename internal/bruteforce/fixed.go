package bruteforce

import "github.com/GustavoRutkowski/cryptowski/internal/cryptowski"

func Fixed(data []byte, keysize uint8, decodeAlg cryptowski.DecodeAlg) (string, []byte, error) {
	bytes := make([]byte, keysize)
	indexes := make([]int, keysize)

	for i := range keysize {
		bytes[i] = KEY_CHARSET[0]
	}

	i := int(keysize) - 1
	for i >= 0 {
		key := string(bytes)
		decoded := decodeAlg(data, key)

		success, err := VerifyResult(decoded)
		if err != nil {
			return "", nil, err
		}

		if success {
			return key, decoded, nil
		}

		for i >= 0 {
			indexes[i]++
			if indexes[i] < len(KEY_CHARSET) {
				bytes[i] = KEY_CHARSET[indexes[i]]
				break
			}

			indexes[i] = 0
			bytes[i] = KEY_CHARSET[0]
			i--
		}
	}

	return "", nil, nil
}
