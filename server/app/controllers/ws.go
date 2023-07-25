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

func SetupWsController(s *socketio.Server) {
	socket = s
}

func MsgAddedEventHandlre(s socketio.Conn, data map[string]string) error {
	// Store msg in db
	err := services.LoadMessageService().CreateMessage(data["from"], data["to"], data["message"])
	if err!= nil {
		return err
	}

	// Prepair response body
	id , err := primitive.ObjectIDFromHex(data["from"])
	var responePayload = models.MessageRead{
		Sender: id,
		Users: []string{data["from"], data["to"]},
		Message: data["message"],
	} 

	// To json
	responseData, err := json.Marshal(responePayload)

	// Send to contact user
	sendToConn(data["to"], string(responseData))

	return nil
}

func HandleConnection(s socketio.Conn) error {
	id := getIdFromRawQuery(s.URL().RawQuery)

	// Add new client
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

func sendToConn(connId string, data string) {
	for _, client := range clients {
		if client.connId == connId {
			client.socketConn.Emit("entry-message", data)
		}
	}
}