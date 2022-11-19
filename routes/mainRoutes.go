package routes

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rohitCodeRed/go_crypto/blockchain"
	"github.com/rohitCodeRed/go_crypto/config"
	"github.com/rohitCodeRed/go_crypto/controllers"
)

func Router(b *blockchain.BlockChain, app *config.AppConfig) http.Handler {
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
	r.Use(CORSMiddleware())

	// routes := e.Group("/api")
	// {
	// 	routes.POST("/login", loginEndpoint)
	// 	routes.POST("/submit", submitEndpoint)
	// 	routes.POST("/read", readEndpoint)
	// }
	r.GET("/", func(c *gin.Context) {
		controllers.Repo.Home(c, b)
		// c.JSON(
		// 	http.StatusOK,
		// 	gin.H{
		// 		"code":    http.StatusOK,
		// 		"message": "Welcome to Crypto currency rocoin Node with address: " + b.GetUuidAddress(),
		// 	},
		// )
	})

	r.POST("/login", controllers.Login)

	r.GET("/mine_block", func(c *gin.Context) {
		controllers.MineBlock(c, b)
	})

	r.GET("/get_chain", func(c *gin.Context) {
		controllers.GetChain(c, b)
	})

	r.GET("/is_valid", func(c *gin.Context) {
		controllers.IsChainValid(c, b)
	})

	r.POST("/add_transaction", func(c *gin.Context) {
		controllers.AddTransaction(c, b)
	})
	r.POST("/update_transaction", func(c *gin.Context) {
		controllers.UpdateTransaction(c, b)
	})

	r.POST("/connect_node", func(c *gin.Context) {
		controllers.ConnectNode(c, b)
	})

	r.GET("/replace_chain", func(c *gin.Context) {
		controllers.Replace_chain(c, b)
	})

	r.GET(("/socket_conn"), func(c *gin.Context) {
		controllers.GetRealTimeData(c, b)
	})

	return r
}

// MIDDELE WARE... to authenticate..
type AuthHeader struct {
	Authorization string `header:"Authorization"`
	ContentType   string `header:"Content-Type"`
}

func Authenticate(b *blockchain.BlockChain) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		h := AuthHeader{}

		if err := c.ShouldBindHeader(&h); err != nil {
			log.Println(err, "header not found..")
			//c.JSON(http.StatusOK, err)
			c.Set("authenticated", false)
		} else {
			log.Println("Header found!..")
			c.Set("authenticated", true)
		}

		log.Println("Node Address: ", b.GetUuidAddress())

		//TODO get userId...

		// Set example variable
		c.Set("userId", "rohit")
		c.Set("node_address", b.GetUuidAddress())

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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
