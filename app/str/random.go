package str

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/typomedia/gitti/app/msg"
)

func Hex() string {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	msg.Check(err)

	return hex.EncodeToString(bytes)
}
