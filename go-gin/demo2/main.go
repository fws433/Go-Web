package main
import (
	"net/http"

	"github.com/gin-gonic/gin"
)
func main(){
	router:=gin.Default()
	//使用gin的query参数形式
	router.GET("/appGet",getParameter)
	router.Run(":10086")

}
//使用query获取参数
func getParameter(c *gin.Context) {
	nickname:=c.Query("nickname")
	passwd:=c.Query("passwd")
	c.String(http.StatusOK,"参数:%s get parameter ok",nickname,passwd)
}
