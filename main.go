package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("provide host address")
		return
	}
	l, err := net.Listen("tcp", os.Args[1])
	defer l.Close()
	if err != nil {
		fmt.Println("error listening port:", err)
		return
	}
	fmt.Println("listening on", os.Args[1])

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("error accepting request:", err)
			return
		}
		go handleConn(conn)
	}

}

func getFile(path string) string {
	// For security reasons so a user can't access something like /etc/passwd
	// All files should be in server directory
	pwd, _ := os.Getwd()
	f := filepath.Join(pwd, path)
	return f
}

func handleConn(conn net.Conn) {
	req, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("error reading request:", err)
		resp := getErrorHTTPResponse([]byte(""), 500)
		conn.Write([]byte(resp))
		return
	}

  fmt.Println(req)
	reqFile := strings.Split(req, " ")[1]
	f := getFile(reqFile)
	if _, err := os.Stat(f); err != nil {
		fmt.Println(err)
		resp := getErrorHTTPResponse([]byte(""), 404)
		conn.Write([]byte(resp))
		return
	}
	contents, err := os.ReadFile(f)
	if err != nil {
		return
	}
	ext := filepath.Ext(f)
	resp := getOKHTTPResponse(contents, ext)
	conn.Write([]byte(resp))

}

func getOKHTTPResponse(body []byte, ext string) (res string) {
	var conType string
	switch ext {
	case ".html":
		conType = "text/html"
	default:
		conType = "text/plain"
	}
	return fmt.Sprintf(
		"HTTP/1.1 200 OK\r\n"+
			"Content-Type: %s\r\n"+
			"Content-Length: %d\r\n"+
			"\r\n"+
			"%s",
		conType,
		len(body),
		string(body),
	)
}
func getErrorHTTPResponse(body []byte, code int) (res string) {
	switch code {
	case 404:
		return fmt.Sprintf(
			"HTTP/1.1 404 Not Found\r\n"+
				"Content-Type: text/plain\r\n"+
				"Content-Length: %d\r\n"+
				"\r\n"+
				"%s",
			len(body),
			body,
		)
	default:
		return fmt.Sprintf(
			"HTTP/1.1 500 InternalServerError\r\n"+
				"Content-Type: text/plain\r\n"+
				"Content-Length: %d\r\n"+
				"\r\n"+
				"%s",
			len(body),
			string(body),
		)
	}

}

