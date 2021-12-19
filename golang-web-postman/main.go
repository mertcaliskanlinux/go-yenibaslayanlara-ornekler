package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func blog_post(c *gin.Context) {
	//url'den gelenler
	no := c.Query("no")
	sayfa := c.Query("sayfa")
	//PostForm => POST Verisinden Gelenler
	isim := c.PostForm("isim")
	mesaj := c.PostForm("mesaj")

	fmt.Printf("no:%s;,sayfa:%s;,isim:%s;,mesaj:%s\n", no, sayfa, isim, mesaj)
}

func main() {
	yönlendirici := gin.Default()
	yönlendirici.POST("/blog_post", blog_post)
	yönlendirici.Run(":8000")

}
