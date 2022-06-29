package model

import "time"

type User struct {
	ID         uint      `gorm:"primaryKey"`
	LoginName  string    `gorm:"column:loginName"`
	CreateTime time.Time `gorm:"column:createTime"`
	UpdateTime time.Time `gorm:"column:updateTime"`
}
