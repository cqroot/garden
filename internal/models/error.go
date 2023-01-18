package models

import (
	"errors"

	"gorm.io/gorm"
)

func IsErrNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
