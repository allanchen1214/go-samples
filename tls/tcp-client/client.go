package main

import (
	"crypto/tls"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)

	conf := &tls.Config{
		InsecureSkipVerify: true, // 忽略证书检查

	}

	conn, err := tls.Dial("tcp", "127.0.0.1:8081", conf)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	n, err := conn.Write([]byte("hello\n"))
	if err != nil {
		log.Println(n, err)
		return
	}

	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		log.Println(n, err)
		return
	}

	println(string(buf[:n]))
}
