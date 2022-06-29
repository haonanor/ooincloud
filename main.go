package main

import (
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"study/cmd"
)

func main() {
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() // 监听并在 0.0.0.0:8080 上启动服务
	cmd.Execute()

}

const StaticFolder = "statics"

func Eject(static fs.FS) {
	embedFS, err := fs.Sub(static, "dist/dist")
	if err != nil {
		// fmt.Println("aaaa")
		// fmt.Printf("无法初始化静态资源, %s", err)
		return
	}
	var walk func(relPath string, d fs.DirEntry, err error) error
	walk = func(relPath string, d fs.DirEntry, err error) error {
		if err != nil {
			panic("无法获取信息")
		}

		if !d.IsDir() {
			fmt.Println(relPath)
			// out, _ := CreatNestedFile(filepath.Join(RelativePath(""), StaticFolder, relPath))
			// defer out.Close()
			// obj, _ := embedFS.Open(relPath)
			// if _, err := io.Copy(out, bufio.NewReader(obj)); err != nil {
			// 	return errors.Errorf("无法写入文件[%s], %s, 跳过...", relPath, err)
			// }
			// // 写入文件
			// out, err := util.CreatNestedFile(filepath.Join(util.RelativePath(""), StaticFolder, relPath))
			// defer out.Close()

			// if err != nil {
			// 	panic("无法创建文件")
			// }

			// obj, _ := embedFS.Open(relPath)
			// if _, err := io.Copy(out, bufio.NewReader(obj)); err != nil {
			// 	return errors.Errorf("无法写入文件[%s], %s, 跳过...", relPath, err)
			// }
		}
		return nil
	}
	err = fs.WalkDir(embedFS, ".", walk)

	fmt.Println(err)
	// var walk func(relPath string, d fs.DirEntry, err error) error
	// walk = func(relPath string, d fs.DirEntry, err error) error {
	// 	if err != nil {
	// 		panic("无法获取信息")
	// 	}

	// 	if !d.IsDir() {
	// 		out, _ := CreatNestedFile(filepath.Join(RelativePath(""), StaticFolder, relPath))
	// 		defer out.Close()
	// 		obj, _ := embedFS.Open(relPath)
	// 		if _, err := io.Copy(out, bufio.NewReader(obj)); err != nil {
	// 			return errors.Errorf("无法写入文件[%s], %s, 跳过...", relPath, err)
	// 		}
	// 		// // 写入文件
	// 		// out, err := util.CreatNestedFile(filepath.Join(util.RelativePath(""), StaticFolder, relPath))
	// 		// defer out.Close()

	// 		// if err != nil {
	// 		// 	panic("无法创建文件")
	// 		// }

	// 		// obj, _ := embedFS.Open(relPath)
	// 		// if _, err := io.Copy(out, bufio.NewReader(obj)); err != nil {
	// 		// 	return errors.Errorf("无法写入文件[%s], %s, 跳过...", relPath, err)
	// 		// }
	// 	}
	// 	return nil
	// }

	// fmt.Println(walk)
	// fmt.Println(embedFs)
	// // util.Log().Info("开始导出内置静态资源...")
	// err = fs.WalkDir(embedFS, ".", walk)
	// if err != nil {
	// 	util.Log().Error("导出内置静态资源遇到错误：%s", err)
	// 	return
	// }
	// util.Log().Info("内置静态资源导出完成")
}
func RelativePath(name string) string {
	if filepath.IsAbs(name) {
		return name
	}

	e, _ := os.Executable()
	return filepath.Join(filepath.Dir(e), name)
}
func CreatNestedFile(path string) (*os.File, error) {
	basePath := filepath.Dir(path)
	if !Exists(basePath) {
		err := os.MkdirAll(basePath, 0700)
		if err != nil {

			return nil, err
		}
	}

	return os.Create(path)
}
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
