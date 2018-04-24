package models

import (
	"github.com/hjkelly/zbbapi/common"
	uuid "github.com/satori/go.uuid"
)

// SafeUUID contains a UUID without blowing up JSON unmarshaling when an ID is missing or improperly formatted. TODO: This may not be necessary if we can use UnsupportedFieldError or UnsupportedValueError.
type SafeUUID string

// NewSafeUUID creates a v4 UUID in a reusable way.
func NewSafeUUID() SafeUUID {
	return SafeUUID(uuid.NewV4().String())
}

// GetValidated returns a normalized UUID if it's valid, otherwise an error.
func (raw SafeUUID) GetValidated() (SafeUUID, error) {
	properUUID, err := uuid.FromString(string(raw))
	if err != nil {
		return "", common.NewValidationError("", common.BadUUIDFormatCode, "Double-check the ID you're trying to reference, because this one doesn't look right. It should be in the format of a UUID.")
	}
	return SafeUUID(properUUID.String()), nil
}
