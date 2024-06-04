package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"strconv"
	"time"
)

func GenerateID(prefix string, length int) string {
	// Get the current timestamp
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)

	// Generate a random number
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err) // handle error appropriately in real code
	}
	randomString := hex.EncodeToString(randomBytes)

	// Combine timestamp and random string
	combined := timestamp + randomString

	// Hashing the combined string
	hasher := md5.New()
	hasher.Write([]byte(combined))
	hash := hex.EncodeToString(hasher.Sum(nil))

	// Get the substring of the hash based on length
	id := prefix + hash[:length]

	return id
}
