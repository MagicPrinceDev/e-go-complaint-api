package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
)

func GenerateID(prefix string, length int) string {
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)

	// Hashing
	hasher := md5.New()
	hasher.Write([]byte(timestamp))
	hash := hex.EncodeToString(hasher.Sum(nil))

	//  Get substring of hash based on length
	id := prefix + hash[:length]

	return id
}
