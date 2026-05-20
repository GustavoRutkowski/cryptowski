package cryptowski

type CryptoAlg func(data []byte, key string) []byte
type EncodeAlg func(data []byte, key string) []byte
type DecodeAlg func(data []byte, key string) []byte

type ICrypto interface {
	Encode(data []byte, key string) []byte
	Decode(data []byte, key string) []byte
}
