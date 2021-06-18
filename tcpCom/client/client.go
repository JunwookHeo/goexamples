package main

import (
	"bytes"
	crand "crypto/rand"
	"io"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/junwookheo/goexamples/tcpCom/serial"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	for {
		n := 1024 + rand.Intn(1024)
		tx := randString(n)
		//log.Printf("tx : %s", tx)
		log.Printf("Send data length : %d, %s ... %s", len(tx), tx[:10], tx[len(tx)-10:])

		sendData(tx)
		time.Sleep(time.Duration(3) * time.Second)
	}
}

func sendData(tx string) {
	conn, err := net.Dial("tcp", ":5032")
	if err != nil {
		log.Fatalf("Failed to connect to the server")
	}
	defer conn.Close()

	btx := serial.Serialize([]byte(tx))

	_, err = io.Copy(conn, bytes.NewReader(btx))
	if err != nil {
		log.Fatalf("Failed to send data to the server : %v", err)
	}
	//conn.Write(btx)
}

func randString(n int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	crand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}

	return string(bytes)
}
