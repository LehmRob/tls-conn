package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"log"
	"net"
)

var (
	certPath = flag.String("cert", "", "Path to certificate")
	keyPath  = flag.String("key", "", "Path to key file")
)

func main() {
	flag.Parse()
	cerr, err := tls.LoadX509KeyPair(*certPath, *keyPath)
	if err != nil {
		log.Fatal(err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cerr},
	}

	sock, err := tls.Listen("tcp", ":8443", config)
	if err != nil {
		log.Fatal(err)
	}
	defer sock.Close()

	for {
		log.Print("Waiting for connection")
		client, err := sock.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(client)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	log.Print("Connection accepted")
	r := bufio.NewReader(conn)
	msg, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println(msg)
	_, err = conn.Write([]byte("Hello World\n"))
	if err != nil {
		log.Fatal(err)
		return
	}
}
