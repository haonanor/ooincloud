package v1

import (
	"study/common"
	"study/core"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	common.Base
}

func (o UserController) Register(c *gin.Context) {
	o.SetContext(c)
	// e := core.ROBOT_CONFIG.Get("app")
	// fmt.Println(e)
	// user := model.User{LoginName: "hahahhaha", CreateTime: time.Now(), UpdateTime: time.Now()}

	// res := core.ROBOT_DB.Create(&user)
	// fmt.Println(res.RowsAffected)
	// fmt.Println(user.ID)
	core.ROBOT_LOGGER.Debug("ahhahahaha")
	// o.Orm.
	o.ERROR(12, "aaaa")
}
