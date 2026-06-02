package bruteforce

import (
	"strings"
	"unicode"
)

type rate float32
const CHARSET_RATE rate = 0.8
const DICTIONARY_RATE rate = 0.5

func validCharset(decoded []byte) rate {
	var bytes uint = 0
	var valid uint = 0

	for _, char := range decoded {
		bytes++
		r := rune(char)
		if unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsPunct(r) || unicode.IsSpace(r) {
			valid++
		}
	}

	return rate(valid) / rate(bytes)
}

var dict, dictErr = DownloadDictionary()

func validWords(decoded []byte) (rate, error) {
	var words uint = 0
	var valid uint = 0

	if dict == nil {
		dict, dictErr = DownloadDictionary()
	}

	wordsList := strings.Fields(string(decoded))

	if dictErr != nil {
		return 0, dictErr
	}

	for _, w := range wordsList {
		words++
		if dict.Has(w) {
			valid++
		}
	}

	return rate(valid) / rate(words), nil
}

// Checks if the decoded text is valid (if it has common characters and known words).
func verifyResult(decoded []byte) (bool, error) {
	if validCharset(decoded) < CHARSET_RATE {
		return false, nil
	}

	validWordsRate, err := validWords(decoded)
	if err != nil {
		return false, err
	}

	if validWordsRate <= DICTIONARY_RATE {
		return false, nil
	}

	return true, nil
}
