package datasource

import (
	"encoding/base64"
	"github.com/gofrs/uuid"
)

func UUIDFromString(uuidString string) uuid.UUID {
	var realUUID uuid.UUID
	uuidBytes, err := base64.StdEncoding.DecodeString(uuidString)
	if err == nil {
		realUUID = uuid.FromBytesOrNil(uuidBytes)
	} else {
		realUUID = uuid.FromStringOrNil(uuidString)
	}
	return realUUID
}

func UUIDBase64FromString(uuidString string) string {
	uuid := uuid.FromStringOrNil(uuidString)
	return base64.StdEncoding.EncodeToString(uuid.Bytes())
}