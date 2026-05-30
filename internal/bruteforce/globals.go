package bruteforce

import "errors"

var ErrKeyNotFound = errors.New("key not found")

// 62 characters in charset.
// Complexity: 62^keysize
const KEY_CHARSET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const KEY_CHARSET_LEN = len(KEY_CHARSET)
