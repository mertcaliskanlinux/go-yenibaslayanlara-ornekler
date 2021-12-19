package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//http://sunucu/kullanici/isim şeklinde çarıldığında cevap veriyor
func Selamla(c *gin.Context) {
	isim := c.Param("isim")
	c.String(http.StatusOK, "Merhaba %s", isim)
}

//http/sunucu/kullanici/isim/sil şeklinde çağrıldığında "mert" için "sil" İŞLEMİ işleniyor
func eylem(c *gin.Context) {
	isim := c.Param("isim")
	eylem := c.Param("eylem")
	mesaj := isim + " İçin uygulancak olan Eylem %s" + eylem

	c.String(http.StatusOK, mesaj)
}

func main() {
	yönlendirici := gin.Default()
	//adress satırından gelen değişkenleri alalım
	yönlendirici.GET("/blog", Blog)
	yönlendirici.GET("/kullanici/:isim", Selamla)
	yönlendirici.GET("/kullanici/:isim/*eylem", eylem)
	yönlendirici.Run(":8000")

}

func Blog(c *gin.Context) {
	//Varsayılan yazı  numarası 1 olsun
	yazi := c.DefaultQuery("yazi", "1(Varsayilan)")
	eylem := c.Query("eylem")
	c.Param("isim")
	c.String(http.StatusOK, "%s Numaralı Yazı için %s işlem yapılacaktır.", yazi, eylem)
}
