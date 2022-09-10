package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
)
import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("index.html")

	//首页
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	//实现文件上传
	r.POST("/uploadFun", func(c *gin.Context) {
		file, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
		} else {
			dst := path.Join("./", file.Filename)
			err := c.SaveUploadedFile(file, dst)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"tips": "file upload failed",
					"msg":  err.Error(),
				})
			}
			c.JSON(http.StatusOK, gin.H{
				"tips": "file upload success",
			})
		}
	})
	//多文件上传
	r.POST("/multiUploadFun", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["f2"]

		for idx, file := range files {
			log.Println(file.Filename)
			//dst := path.Join("./multiFile", file.Filename)
			dst := fmt.Sprintf("./multiFile/%d_%s", idx, file.Filename)
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files uploaded", len(files)),
		})
	})

	r.Run(":9000")
}
