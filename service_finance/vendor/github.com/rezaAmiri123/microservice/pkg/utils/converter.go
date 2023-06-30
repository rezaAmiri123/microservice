package utils

import (
	"encoding/base64"
	"fmt"

	"github.com/google/uuid"
)

func ConvertBase64ToUUID(value []byte) (uuid.UUID, error) {
	strValue := string(value)
	data, err := base64.StdEncoding.DecodeString(strValue)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("cannot decode: %w", err)
	}
	valueByte, err := uuid.FromBytes(data)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("cannot decode: %w", err)
	}
	return valueByte, nil
}
