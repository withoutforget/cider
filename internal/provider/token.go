package provider

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

type TokenProvider struct {
	tokenLength int
	tokenLoops  int
	salt        []byte
}

func NewTokenProvider() *TokenProvider {
	return &TokenProvider{
		tokenLength: 10,
		tokenLoops:  2,
		salt:        genSalt(10),
	}
}

func genSalt(len int) []byte {
	data := make([]byte, len)
	if _, err := rand.Read(data); err != nil {
		return nil
	}
	return data
}

func (p *TokenProvider) Provide(data []byte) string {
	ret := make([]byte, len(data), int(p.tokenLength)+len(data))
	copy(ret, data)
	ret = append(ret, p.salt...)

	res := sha256.Sum256(ret)

	return hex.EncodeToString(res[:])
}
