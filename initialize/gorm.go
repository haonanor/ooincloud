package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Gorm() *gorm.DB {
	return GormMysql()
}

func GormMysql() *gorm.DB {
	dsn := "root:1fe8890e5084042e@tcp(47.105.88.89:3306)/gitlike?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "oo_", // table name prefix, table for `User` would be `t_users`
			SingularTable: true,
		},
	})
	if err != nil {
		panic("数据库链接失败")
	}
	return db
}
