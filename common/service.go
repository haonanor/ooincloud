package common

import "gorm.io/gorm"

type Service struct {
	Orm *gorm.DB
}
