package controllers

import (
	"encoding/json"
	"log"
	"net/url"

	socketio "github.com/googollee/go-socket.io"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"chat/app/models"
	"chat/app/services"
)

type client struct {
	connId 		string
	userId 		string
	socketConn 	socketio.Conn
}

var clients []*client

var socket *socketio.Server

func MsgAddedEventHandlre(s socketio.Conn, data map[string]string) {
	log.Println(data)
	err := services.LoadMessageService().CreateMessage(data["from"], data["to"], data["message"])
	if err!= nil {
	}

	id , err := primitive.ObjectIDFromHex(data["from"])
	var responePayload models.MessageRead
	responePayload.Sender = id
	responePayload.Users = []string{data["from"], data["to"]}
	responePayload.Message = data["message"]

	responseData, err := json.Marshal(responePayload)

	for _, client := range clients {
		if client.connId == data["to"] {
			client.socketConn.Emit("entry-message", string(responseData))
		}
	}

}

func SetupWsController(s *socketio.Server) {
	socket = s
}

func HandleConnection(s socketio.Conn) error {
	s.SetContext("")
	log.Println("connected:", s.ID())

	id := getIdFromRawQuery(s.URL().RawQuery)

	clients = append(clients, &client{
		userId: id,
		connId: s.ID(),
		socketConn: s,
	})

	return nil
}

func getIdFromRawQuery(s string) string {
	values, _ := url.ParseQuery(s)
	return values.Get("id")
}

func HandleAddUserEvent(s socketio.Conn, _ string) {
	socket.BroadcastToNamespace("/", "user_added")
}

func HandleByeEvent(s socketio.Conn) string {
	last := s.Context().(string)
	s.Emit("bye", last)
	s.Close()
	return last
}

func HandleErr(_ socketio.Conn, e error) {
	log.Println("meet error:", e)
}

func HandleDisconnection(s socketio.Conn, msg string) {
	log.Println("closed", msg)
}