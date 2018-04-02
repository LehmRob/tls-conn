package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"flag"
	"io/ioutil"
	"log"
)

var certPath = flag.String("cert", "", "Path to certificate")

func readCert(path string) (*x509.CertPool, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(content)
	if !ok {
		return nil, errors.New("Can't read certificate")
	}

	return roots, nil
}

func main() {
	flag.Parse()
	roots, err := readCert(*certPath)
	if err != nil {
		log.Fatal(err)
	}

	conf := &tls.Config{RootCAs: roots}

	conn, err := tls.Dial("tcp", "localhost:8443", conf)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("Hello_Friend\n"))
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(conn)
	msg, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Received from server: %s\n", msg)

}
