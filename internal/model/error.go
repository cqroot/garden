package model

import (
	"errors"

	"gorm.io/gorm"
)

func (m Model) IsErrNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
