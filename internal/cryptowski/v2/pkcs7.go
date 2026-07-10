package v2

import "bytes"

type pkcs7 struct {}

func (p pkcs7) padding(data []byte) []byte {
	paddingLength := BLOCK_SIZE_BYTES - (len(data) % BLOCK_SIZE_BYTES)

	if paddingLength == 0 {
		paddingLength = BLOCK_SIZE_BYTES
	}

	pad := bytes.Repeat([]byte{byte(paddingLength)}, paddingLength)
	return append(data, pad...)
}

func (p pkcs7) unpadding(data []byte) []byte {
	lenData := len(data)

	if lenData % BLOCK_SIZE_BYTES != 0 {
		return data
	}

	paddingLength := int(data[lenData - 1])

	if paddingLength <= 0 || paddingLength > BLOCK_SIZE_BYTES {
		return data
	}

	// Validate padding bytes
	for i := lenData - paddingLength; i < lenData; i++ {
		if int(data[i]) != paddingLength {
			return data
		}
	}

	return data[:lenData - paddingLength]
}
