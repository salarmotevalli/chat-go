package controllers

import (
	// "fmt"
	"log"

	socketio "github.com/googollee/go-socket.io"
)

var socket *socketio.Server

func SetupWsController(s *socketio.Server) {
	socket = s
}

func HandleConnection(s socketio.Conn) error {
	s.SetContext("")
	log.Println("connected:", s.ID())

	return nil
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

func HandleDisconnection(_ socketio.Conn, msg string) {
	log.Println("closed", msg)
}