package common

import (
	"github.com/gin-gonic/gin"
)

type Base struct {
	Context *gin.Context
}

func (o *Base) SetContext(c *gin.Context) *Base {
	o.Context = c
	return o
}

// func (o *Base) SetOrm() *Base {
// 	dsn := "root:1fe8890e5084042e@tcp(47.105.88.89:3306)/gitlike?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
// 		NamingStrategy: schema.NamingStrategy{
// 			TablePrefix:   "oo_", // table name prefix, table for `User` would be `t_users`
// 			SingularTable: true,
// 		},
// 	})
// 	if err != nil {
// 		panic("数据库链接失败")
// 	}
// 	o.Orm = db
// 	return o
// }

func (o Base) OK(data interface{}, message string) {
	Ok(o.Context, data, message)
}

func (o Base) ERROR(code int32, message string) {
	Error(o.Context, code, message)
}
