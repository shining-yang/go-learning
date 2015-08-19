package main
import (
	"fmt"
	"bytes"
	"os"
	"io"
	"net"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], " host")
		os.Exit(0)
	}

	service := os.Args[1]

	fmt.Println("Dial up a tcp connection")
	conn, err := net.Dial("tcp", service)
	checkError(err)

	fmt.Println("Send HTTP HEAD request")
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	fmt.Println("Get the response from server")
	response, err := readFull(conn)
	checkError(err)

	fmt.Println(string(response))
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func readFull(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}
	}

	return result.Bytes(), nil
}

