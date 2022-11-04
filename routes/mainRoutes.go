package routes

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rohitCodeRed/go_crypto/blockchain"
	"github.com/rohitCodeRed/go_crypto/controllers"
)

func Router() http.Handler {
	r := gin.New()

	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	r.Use(Authenticate())

	// routes := e.Group("/api")
	// {
	// 	routes.POST("/login", loginEndpoint)
	// 	routes.POST("/submit", submitEndpoint)
	// 	routes.POST("/read", readEndpoint)
	// }
	r.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":    http.StatusOK,
				"message": "Welcome to Crypto currency rocoin Node with address: " + blockchain.GetUuidAddress(),
			},
		)
	})

	r.POST("/login", controllers.Login)

	r.POST("/mine_block", controllers.MineBlock)

	r.GET("/get_chain", controllers.GetChain)

	r.GET("/is_valid", controllers.IsChainValid)

	r.GET("/add_transaction", controllers.AddTransaction)

	r.POST("/connect_node", controllers.ConnectNode)

	r.GET("/replace_chain", controllers.Replace_chain)

	return r
}

// MIDDELE WARE... to authenticate..
type AuthHeader struct {
	Authorization string `header:"Authorization"`
	ContentType   string `header:"Content-Type"`
}

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		h := AuthHeader{}

		if err := c.ShouldBindHeader(&h); err != nil {
			log.Println(err)
			//c.JSON(http.StatusOK, err)
		}

		//TODO get userId...

		// Set example variable
		c.Set("userId", "12345")
		c.Set("node_address", blockchain.GetUuidAddress())

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}

}
