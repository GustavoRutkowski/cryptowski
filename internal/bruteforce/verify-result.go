package bruteforce

import (
	"strings"
	"unicode"
)

func validCharset(decoded []byte) Rate {
	var bytes uint = 0
	var valid uint = 0

	for _, char := range decoded {
		bytes++
		r := rune(char)
		if unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsPunct(r) {
			valid++
		}
	}

	return Rate(bytes) / Rate(valid)
}

func validWords(decoded []byte) (Rate, error) {
	var words uint = 0
	var valid uint = 0

	wordsList := strings.Split(string(decoded), " ")
	dictionary, err := DownloadDictionary()

	if err != nil {
		return 0, err
	}

	for _, w := range wordsList {
		words++
		if dictionary.Has(w) {
			valid++
		}
	}

	return Rate(words) / Rate(valid), nil
}

func VerifyResult(decoded []byte) (bool, error) {
	if validCharset(decoded) < CHARSET_RATE {
		return false, nil
	}

	validWordsRate, err := validWords(decoded)
	if err != nil {
		return false, err
	}

	if validWordsRate < DICTIONARY_RATE {
		return false, nil
	}

	return true, nil
}
