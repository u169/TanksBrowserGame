package utils

import (
	"crypto/rand"
	"fmt"
)

func ObjectId(busied chan string) string {
	oId := genRandomString()
	for !isAvailableId(oId, busied) {
		oId = genRandomString()
	}
	return oId
}

func genRandomString() string {
	b := make([]byte, 8)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func isAvailableId(s string, busied chan string) bool {
	for b := range busied {
		if b == s {
			return false
		}
	}
	return true
}
