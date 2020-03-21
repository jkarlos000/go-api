package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	server, err := socketio.NewServer(nil)

	if err != nil {
		log.Fatalln(err)
	}

	// sockets
	server.OnConnect("/", func(s socketio.Conn) error {
		//s.SetContext("")
		fmt.Println("connected:", s.ID())
		s.Join("chat_room")
		return nil
	})
	server.OnEvent("/", "chat message", func(s socketio.Conn, msg string) {
		server.BroadcastToRoom(s.Namespace(), "chat_room", "chat message", msg)
	})

	go server.Serve()
	defer server.Close()

	//Return Front End

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Println("Server on port: 5000")
	log.Fatalln(http.ListenAndServe(":5000", nil))
}
