package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"

	"chat/app/controllers"
)

func Setup(engine *gin.Engine, socket *socketio.Server)  {
	setupWsRoutes(socket)	
	
	engine.Use(corsMiddleware("http://localhost:3000"))
	
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

func setupWsRoutes(socket *socketio.Server) {
	controllers.SetupWsController(socket)
	
	socket.OnEvent("/", "msg-added", controllers.MsgAddedEventHandlre)
	
	socket.OnConnect("/", controllers.HandleConnection)
	
	// socket.OnEvent("/", "notice", controllers.HandleNoticeEvent)

	// socket.OnEvent("/chat", "msg", controllers.HandleMsgEvent)

	// socket.OnEvent("/", "bye", controllers.HandleByeEvent)

	socket.OnError("/", controllers.HandleErr)

	socket.OnDisconnect("/", controllers.HandleDisconnection)
}
