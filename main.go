package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("sunucu başlatılamadı", err)
	}
	defer listener.Close()

	fmt.Println("sunucu 8080 portunda dinleniyor")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("yeni istemci baglanamadı")
			continue
		}

		fmt.Println("yeni bir istemci baglandı")
		conn.Write([]byte("selam tcp ben yeni istemci\n"))
		conn.Close()
	}
}
