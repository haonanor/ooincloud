package v1

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"study/common"
	"study/core"
	"study/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type FileController struct {
	common.Base
}

func (o FileController) GetA(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	core.ROBOT_CACHE.Set("upload:1", []byte(strconv.Itoa(rand.Intn(10000))))
}
func (o FileController) GetB(c *gin.Context) {
	a, _ := core.ROBOT_CACHE.Get("upload:1")
	fmt.Printf("cast.ToString(a): %v\n", cast.ToString(a))
}
func (o FileController) Post(c *gin.Context) {
	s := c.Request.Header.Get("Upload-Length")
	if s == "" {
		// uploadPath := "uploads/"
		uploadPath := core.ROBOT_CONFIG.Get("upload.path").(string)
		file, _ := c.FormFile("file")
		core.ROBOT_LOGGER.Info(file.Filename)
		c.SaveUploadedFile(file, uploadPath+utils.CreateFileKey()+"."+utils.GetFileExt(file.Filename))
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	} else {
		c.Header("Content-Type", "text/plain;")
		c.String(http.StatusOK, utils.CreateFileKey())
	}

}

func (o FileController) Head(c *gin.Context) {
	id := c.Query("patch")
	b, err := core.ROBOT_CACHE.Get("upload:" + id)
	if err != nil {
		c.String(http.StatusInternalServerError, "重启失败")
		return
	}

	c.Header("Upload-Offset", cast.ToString(b))
	c.String(http.StatusOK, "继续上传")
}

func (o FileController) Patch(c *gin.Context) {
	o.SetContext(c)
	var err error
	uploadPath := core.ROBOT_CONFIG.Get("upload.path").(string)
	fmt.Printf("uploadPath: %v\n", uploadPath)
	id := c.Query("patch")
	offsetStr := c.Request.Header.Get("Upload-Offset")
	// if offsetStr == "200000000" {
	// 	c.String(http.StatusInternalServerError, "上传失败")
	// 	return
	// }
	uploadFileName := c.Request.Header.Get("Upload-Name")

	outputFile := uploadPath + id + "." + utils.GetFileExt(uploadFileName)

	bdoy, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		core.ROBOT_LOGGER.Error("body读取失败")
		c.String(http.StatusInternalServerError, "上传失败")
		return
	}

	file, err := OpenFile(outputFile)
	if err != nil {
		core.ROBOT_LOGGER.Error("目标文件打开失败")
		c.String(http.StatusInternalServerError, "目标文件打开失败")
		return
	}
	defer file.Close()
	offset, _ := strconv.ParseInt(offsetStr, 10, 64)

	file.WriteAt(bdoy, offset)
	step := core.ROBOT_CONFIG.Get("upload.chunkSize").(int)
	core.ROBOT_CACHE.Set("upload:"+id, []byte(cast.ToString(offset+cast.ToInt64(step))))
	o.OK(nil, "上传成功")
}
func OpenFile(filename string) (*os.File, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return os.Create(filename) //创建文件
	}
	return os.OpenFile(filename, os.O_RDWR, 0666) //打开文件
}

func (o FileController) Delete(c *gin.Context) {
	c.String(http.StatusOK, "成功")
}
