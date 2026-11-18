package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
)

func parseCommand(reader *bufio.Reader) ([]string, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	line = strings.TrimSpace(line)
	if line == "" {
		return nil, nil
	}

	if line[0] != '*' {
		return strings.Fields(line), nil
	}

	count, err := strconv.Atoi(line[1:])
	if err != nil || count <= 0 {
		return nil, fmt.Errorf("gecersiz dizi uzunlugu")
	}

	args := make([]string, 0, count)

	for i := 0; i < count; i++ {
		lenLine, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		lenLine = strings.TrimSpace(lenLine)

		if len(lenLine) < 2 || lenLine[0] != '$' {
			return nil, fmt.Errorf("beklenen bulk string ($) bulunamadi")
		}

		strLen, err := strconv.Atoi(lenLine[1:])
		if err != nil || strLen < 0 {
			return nil, fmt.Errorf("gecersiz string uzunlugu")
		}

		buf := make([]byte, strLen+2)
		if _, err := io.ReadFull(reader, buf); err != nil {
			return nil, err
		}

		args = append(args, string(buf[:strLen]))
	}

	return args, nil
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Yeni istemci bağlandı:", conn.RemoteAddr())

	reader := bufio.NewReader(conn)

	for {
		args, err := parseCommand(reader)
		if err != nil {
			if err == io.EOF {
				fmt.Println("İstemci ayrıldı:", conn.RemoteAddr())
			} else {
				fmt.Println("Protokol/okuma hatası:", err)
				conn.Write([]byte("-ERR " + err.Error() + "\r\n"))
			}
			break
		}

		if len(args) == 0 {
			continue
		}

		fmt.Printf("[%s] Ayrıştırılan Argümanlar: %q\n", conn.RemoteAddr(), args)

		cmd := strings.ToUpper(args[0])

		switch cmd {
		case "PING":
			conn.Write([]byte("+PONG\r\n"))
		case "QUIT":
			conn.Write([]byte("+OK\r\n"))
			return
		default:
			conn.Write([]byte("-ERR bilinmeyen komut '" + cmd + "'\r\n"))
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println("6379 portu dinlenemedi:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Mini-Redis sunucusu 6379 portunda dinleniyor...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Bağlantı kabul hatası:", err)
			continue
		}

		go handleConnection(conn)
	}
}
