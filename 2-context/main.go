package main
import(
	"net/http"
	"gin2"
)
func main(){
	r := gin2.New()
	r.GET("/",func(c *gin2.Context){
		c.HTML(http.StatusOK,"<h1>Hello Gin2</h1>")
	})

	r.GET("/hello",func(c *gin2.Context){
		c.String(http.StatusOK, "hello %s, you're at %s\n",c.Query("name"),c.Path)
	})

	r.POST("/login", func(c *gin2.Context){
		c.JSON(http.StatusOK,gin2.H{
			"username":c.PostForm("username"),
			"password":c.PostForm("password"),
		})
	})
	r.Run(":9999")
}

