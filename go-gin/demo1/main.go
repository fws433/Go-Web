package main
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 处理带参数的path-GET
func getParameter(c *gin.Context) {
	// 回复一个200 OK，获取传入的参数
	nickname := c.Param("nickname")
	passwd := c.Param("passwd")
	c.String(http.StatusOK, "参数:%s %s  get parameter OK", nickname, passwd)
}

// 处理带参数的path-POST
func postParameter(c *gin.Context) {
	// 回复一个200 OK，获取传入的参数
	nickname := c.Param("nickname")
	passwd := c.Param("passwd")
	c.String(http.StatusOK, "参数:%s %s  post parameter OK", nickname, passwd)
}

// 注意':'和'*'的区别
func optionalParameter(c *gin.Context) {
	// 回复一个200 OK，获取传入的参数
	nickname := c.Param("nickname")
	passwd := c.Param("passwd")
	c.String(http.StatusOK, "参数:%s %s  optional parameter OK", nickname, passwd)
}

func main() {
	router := gin.Default()
	// 注意':'必须要匹配,'*'选择匹配,即存在就匹配,否则可以不考虑
	router.GET("/appGet/:nickname/:passwd", getParameter)
	router.POST("/appPost/:nickname/:passwd", postParameter)
	router.GET("/optional/:nickname/*passwd", optionalParameter)

	router.Run(":10086")
}
