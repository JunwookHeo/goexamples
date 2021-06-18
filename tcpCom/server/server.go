package main

import (
	"io"
	"io/ioutil"
	"log"
	"net"

	"github.com/junwookheo/goexamples/tcpCom/serial"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	l, err := net.Listen("tcp", ":5032")
	if err != nil {
		log.Fatalf("fail to bind address to 5032 with err : %v", err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Fail to accept with err : %v", err)
			continue
		}
		go handlerConn(conn)
	}
}

func handlerConn(conn net.Conn) {
	//rxbuf := make([]byte, 32)
	for {
		//n, err := conn.Read(rxbuf)
		rxbuf, err := ioutil.ReadAll(conn)
		n := len(rxbuf)
		if nil != err {
			if io.EOF == err {
				log.Printf("connection closed : %v", conn.RemoteAddr().String())
				return
			}
			log.Printf("Fail to receive data with err : %v", err)
			return
		}

		if n > 0 {
			brx := serial.Deserialize(rxbuf[:n])
			rx := string(brx)
			log.Printf("Received data length : %d, %s ... %s", len(rx), rx[:10], rx[len(rx)-10:])
		}
	}
}
