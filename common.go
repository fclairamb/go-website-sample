package main

import (
	"crypto/sha1"
	"encoding/hex"
)

func HashSHA1(value string) string {
	h := sha1.New()
	h.Write([]byte(value))
	return hex.EncodeToString(h.Sum(nil))
}
