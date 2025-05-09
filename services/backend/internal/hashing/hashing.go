package hashing

import (
	"crypto/sha256"
	"fmt"
)

func SHA256(in string) string {
	h := sha256.New()
	h.Write([]byte(in))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
