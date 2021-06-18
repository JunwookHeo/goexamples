package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("Accept error : ", err)
		}

		go proxy(conn)
	}
}

func proxy(conn net.Conn) {
	defer conn.Close()

	remote, err := net.Dial("tcp", "gophercon.com:443")
	if err != nil {
		log.Println("Connecting gophercon.com error : ", err)
		return
	}

	defer remote.Close()

	go io.Copy(remote, conn)
	io.Copy(conn, remote)
}
func copyTostderr(conn net.Conn) {
	defer conn.Close()

	for {
		conn.SetReadDeadline(time.Now().Add(3 * time.Second))
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Printf("Finished with err = %v", err)
			return
		}
		os.Stderr.Write(buf[:n])
	}
}
