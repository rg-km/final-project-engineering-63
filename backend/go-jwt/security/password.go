package security

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hex.EncodeToString(hash.Sum(nil)))
}
