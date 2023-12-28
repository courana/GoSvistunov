package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
)

type ClockServer struct {
	Address  string
	TimeZone string
}

func handleClockServer(server ClockServer, wg *sync.WaitGroup) {
	defer wg.Done()

	conn, err := net.Dial("tcp", server.Address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	io.Copy(os.Stdout, conn)
}

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		log.Fatal("Usage: clockwall <server1:port1:timezone1> <server2:port2:timezone2> ...")
	}

	var servers []ClockServer
	for _, arg := range args {
		parts := strings.Split(arg, ":")
		if len(parts) != 3 {
			log.Fatal("Неправильный формат аргументов. Используйте <server:port:timezone>")
		}
		port, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal("Неправильное число порта")
		}

		servers = append(servers, ClockServer{
			Address:  parts[0] + ":" + strconv.Itoa(port),
			TimeZone: parts[2],
		})
	}

	var wg sync.WaitGroup

	for _, server := range servers {
		wg.Add(1)
		go handleClockServer(server, &wg)
	}

	wg.Wait()
}
