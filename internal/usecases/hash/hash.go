package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"hash"
)

var (
	hasher   hash.Hash
	hashSalt = "128efhh13s"
)

func init() {
	hasher = sha256.New()
}

func GetHash(s string) string {
	hasher.Reset()
	hasher.Write([]byte(s))
	return hex.EncodeToString(hasher.Sum(nil))
}
