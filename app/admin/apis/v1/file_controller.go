package v1

import (
	"fmt"
	"path"
	"study/common"

	"github.com/gin-gonic/gin"
)

type FileController struct {
	common.Base
}

func (o FileController) Upload(c *gin.Context) {
	o.SetContext(c)

	fh, err := c.FormFile("file")
	fmt.Printf("fh: %v\n", fh)
	if err == nil {
		s := path.Join("./uploads", fh.Filename)
		err2 := c.SaveUploadedFile(fh, s)
		if err2 == nil {
			o.OK(nil, "上传成功")
			return
		} else {
			o.ERROR(300001, err2.Error())
			return
		}

	}

	o.ERROR(300001, "上传失败")
}
