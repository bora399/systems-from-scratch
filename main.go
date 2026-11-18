package main

import (
	"fmt"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("veri okuma hatası:", err)
		return
	}

	request := string(buffer[:n])
	lines := strings.Split(request, "\r\n")
	parts := strings.Fields(lines[0])

	if len(parts) < 3 {
		return
	}

	method := parts[0]
	path := parts[1]
	fmt.Printf("Metot: %s | Yol: %s\n", method, path)

	var statusLine, contentType, body string

	switch {
	case method == "GET" && path == "/":
		statusLine = "HTTP/1.1 200 OK"
		contentType = "text/plain; charset=utf-8"
		body = "Ana Sayfaya Hos Geldiniz!"

	case method == "GET" && path == "/health":
		statusLine = "HTTP/1.1 200 OK"
		contentType = "application/json"
		body = `{"status":"ok","server":"systems-from-scratch"}`

	default:
		statusLine = "HTTP/1.1 404 Not Found"
		contentType = "text/plain; charset=utf-8"
		body = "404 Aradiginiz Sayfa Bulunamadi"
	}

	response := fmt.Sprintf(
		"%s\r\nContent-Type: %s\r\nContent-Length: %d\r\nConnection: close\r\n\r\n%s",
		statusLine,
		contentType,
		len(body),
		body,
	)

	conn.Write([]byte(response))
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("sunucu baslatılamadı:", err)
		return
	}
	defer listener.Close()

	fmt.Println("8080 portunda sunucu dinleniyor...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("bağlantı kabul error : ", err)
			continue
		}

		go handleConnection(conn)
	}
}
