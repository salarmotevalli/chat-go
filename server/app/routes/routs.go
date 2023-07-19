package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"

	"chat/app/controllers"
)

func Setup(engine *gin.Engine, socket *socketio.Server)  {
	
	socket.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("connected:", s.ID())
		return nil
	})

	socket.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		log.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})

	socket.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})

	socket.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	socket.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	socket.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason)
	})


	
	engine.Use(corsMiddleware("http://localhost:3000"))
	
	engine.GET("/socket.io", gin.WrapH(socket))
	engine.POST("/socket.io", gin.WrapH(socket))
	engine.StaticFS("/public", http.Dir("./asset"))


	api := engine.Group("/api")
	setMessageRoutes(api)
	setAuthRoutes(api)
}

func corsMiddleware(allowOrigin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Request.Header.Del("Origin")

		c.Next()
	}
}

func setMessageRoutes(api *gin.RouterGroup) {
	// Set prefix
	msg := api.Group("/messages")

	msg.POST("/addmsg", controllers.AddMessage)
	msg.POST("/getmsg", controllers.GetMessage)
}
