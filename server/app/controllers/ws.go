package controllers

import (
	// "context"
	"log"
	"strings"

	socketio "github.com/googollee/go-socket.io"

	// "chat/app/db"
)

var clients = make(map[string]*socketio.Conn)

var socket *socketio.Server

func MsgAddedEventHandlre(s socketio.Conn, msg any) {
	log.Println(msg)
}

func SetupWsController(s *socketio.Server) {
	socket = s
}

func HandleConnection(s socketio.Conn) error {
	s.SetContext("")
	log.Println("connected:", s.ID())

	clients[s.ID()] = &s

	id := getIdFromRawQuery(s.URL().RawQuery)

	return nil
}

func getIdFromRawQuery(s string) string {
    return strings.Replace(s, "id=", "", 1)
}

func HandleNoticeEvent(s socketio.Conn, msg string) {
	log.Println("notice:", msg)
	s.Emit("reply", "have "+msg)
}

func HandleMsgEvent(s socketio.Conn, msg string) string {
	s.SetContext(msg)
	return "recv " + msg
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
	delete(clients, s.ID())
}