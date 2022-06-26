package security

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	sha := hex.EncodeToString(hash.Sum(nil))
	return fmt.Sprint(sha)
}
