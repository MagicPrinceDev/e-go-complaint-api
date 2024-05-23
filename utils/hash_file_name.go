package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"path/filepath"
	"time"
)

func HashFileName(filename string) string {
	ext := filepath.Ext(filename)
	name := filename[:len(filename)-len(ext)]
	uniqueID := time.Now().UnixNano() // Gunakan timestamp sebagai entitas unik
	hash := sha256.Sum256([]byte(name + fmt.Sprint(uniqueID)))
	hashedFilename := hex.EncodeToString(hash[:]) + ext
	return hashedFilename
}
