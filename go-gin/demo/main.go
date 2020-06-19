package main
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 注册一个默认的路由器
	router := gin.Default()
	// 最基本的用法
	router.GET("/appGet", func(c *gin.Context) {
		// 回复一个200OK,在client的http-get的resp的body中获取数据
		c.String(http.StatusOK, "appGet OK")
	})
	router.POST("/appPost", func(c *gin.Context) {
		// 回复一个200 OK, 在client的http-post的resp的body中获取数据
		c.String(http.StatusOK, "appPost OK")
	})
	router.PUT("/appPut", func(c *gin.Context) {})
	router.DELETE("/appDelete", func(c *gin.Context) {})
	router.PATCH("/appPatch", func(c *gin.Context) {})
	router.HEAD("/appHead", func(c *gin.Context) {})
	router.OPTIONS("/appOptions", func(c *gin.Context) {})

	// router.Run() // 不传参数，默认是8080端口
	// 绑定端口是 8000 ,注意是要加:"
	router.Run(":10086")
}
