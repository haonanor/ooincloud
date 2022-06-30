package v1

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"study/common"
	"study/utils"

	"github.com/gin-gonic/gin"
)

type FileController struct {
	common.Base
}

func (o FileController) Post(c *gin.Context) {
	c.Header("Content-Type", "text/plain;")
	c.String(http.StatusOK, "12351")
}

func (o FileController) Patch(c *gin.Context) {
	o.SetContext(c)
	// pathData := "uploads"

	tmpPath := "tmp/"

	id := c.Query("patch")

	if is, _ := utils.DirExists(tmpPath + id); !is {
		os.Mkdir(tmpPath+id, 0777)
	}

	// fh, err := c.FormFile("file")
	// fmt.Printf("fh: %v\n", fh)
	// fmt.Println(err)
	b, err2 := ioutil.ReadAll(c.Request.Body)
	offset := c.Request.Header.Get("Upload-Offset")
	leng := c.Request.Header.Get("Upload-Length")
	// o.OK(nil, "上传成功")
	if err2 == nil {
		s := path.Join(tmpPath+id, offset+".chunk")
		err := ioutil.WriteFile(s, b, 0777)
		if err != nil {
			o.ERROR(300001, err.Error())
			return
		}
		i, _ := strconv.Atoi(leng)
		i2, _ := strconv.Atoi(offset)
		if (i - i2) <= 500000 {
			var files []string

			err := filepath.Walk(tmpPath+id, func(path string, info os.FileInfo, err error) error {
				files = append(files, path)
				return nil
			})
			if err != nil {
				panic(err)
			}
			outputFile := "uploads/abc.jpeg"
			f, _ := OpenFile(outputFile)
			// defer f.Close()
			for _, file := range files {
				b2, _ := ioutil.ReadFile(file)
				f.Write(b2)
				fmt.Printf("file: %v\n", file)
			}

		}

		o.OK(nil, "上传成功")
		return
		// err3 := c.SaveUploadedFile(b, s)
		// if err3 == nil {
		// 	o.OK(nil, "上传成功")
		// 	return
		// } else {
		// 	o.ERROR(300001, err2.Error())
		// 	return
		// }

	}

	o.ERROR(300001, "上传失败")
}
func OpenFile(filename string) (*os.File, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("文件不存在")
		return os.Create(filename) //创建文件
	}
	fmt.Println("文件存在")
	return os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
}
