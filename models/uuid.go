package models

import (
	"github.com/hjkelly/zbbapi/common"
	uuid "github.com/satori/go.uuid"
)

type SafeUUID string

func NewSafeUUID() SafeUUID {
	return SafeUUID(uuid.NewV4().String())
}

func (raw SafeUUID) GetValidated() (SafeUUID, error) {
	properUUID, err := uuid.FromString(string(raw))
	if err != nil {
		return "", common.NewValidationError("", common.BadUUIDFormatCode, "Double-check the ID you're trying to reference, because this one doesn't look right. It should be in the format of a UUID.")
	}
	return SafeUUID(properUUID.String()), nil
}
