package utils

import (
	"crypto/rand"
	"math/big"
)

func GenerateOTP(length int) string {
	const charset = "0123456789"
	otp := make([]byte, length)
	for i := range otp {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			panic(err)
		}
		otp[i] = charset[num.Int64()]
	}
	return string(otp)
}
