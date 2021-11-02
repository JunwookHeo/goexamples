package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	defer ws.Close()

	if err != nil {
		log.Println("endpoint", err)
		return
	}

	log.Println("Client Connected!!!")
	reader(ws)
}

func reader(conn *websocket.Conn) {
	for {
		mt, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("reader", err)
			return
		}

		log.Println(string(p))

		if err := conn.WriteMessage(mt, p); err != nil {
			log.Println("writer", err)
			return
		}
	}
}

func main() {
	//http.HandleFunc("/", homePage)
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/ws", wsEndpoint)

	log.Println("Listening on port 8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Println(err)
	}
}
