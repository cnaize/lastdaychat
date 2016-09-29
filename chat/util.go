package chat

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRoomId() string {
	return GenerateString(32)
}

func GenerateString(size uint32) string {
	rb := make([]byte, size)
	if _, err := rand.Read(rb); err != nil {
		return "Generating string error"
	}

	return base64.URLEncoding.EncodeToString(rb)
}
