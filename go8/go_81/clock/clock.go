package main

import (
	"flag"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

var timeZone string

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		loc, err := time.LoadLocation(timeZone)
		if err != nil {
			log.Fatal(err)
		}
		currentTime := time.Now().In(loc).Format("15:04:05\n")

		_, err = io.WriteString(c, currentTime)
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	flag.StringVar(&timeZone, "timezone", "UTC", "определённая временная зона для сервера часов")
	port := flag.Int("port", 8000, "определённый порт для сервера часов")
	flag.Parse()

	listener, err := net.Listen("tcp", "localhost:"+strconv.Itoa(*port))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Printf("Сервер часов работает на порту %d в временной зоне %s\n", *port, timeZone)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
