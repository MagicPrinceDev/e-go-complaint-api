package utils

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

func DecodePayload(token string) (map[string]interface{}, error) {
	// Token is in format "header.payload.signature", we only need the payload
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid JWT token format")
	}

	// Decode base64 encoded payload
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}

	// Unmarshal payload into a map
	var payloadMap map[string]interface{}
	err = json.Unmarshal(payload, &payloadMap)
	if err != nil {
		return nil, err
	}

	return payloadMap, nil
}
