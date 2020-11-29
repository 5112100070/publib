package encoding

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"strings"
)

type HMACAuthData struct {
	Method string
	Date   int64
	Path   string
}

func (this HMACAuthData) GenerateHMACHash(secret string) string {
	key := []byte(secret)
	h := hmac.New(sha1.New, key)

	message := fmt.Sprintf(
		"%s\n%s\n%d",
		this.Method,
		this.Path,
		this.Date,
	)
	h.Write([]byte(message))

	hash := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return strings.Replace(hash, " ", "-", -1)
}
