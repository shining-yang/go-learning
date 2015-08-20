package main
import (
	"fmt"
	"os"
	"net"
	"io/ioutil"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]

	fmt.Println("Resolve TCP address")
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	fmt.Println("Dial up TCP")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	fmt.Println("Send message to server")
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	fmt.Println("Getting response")
	response, err := ioutil.ReadAll(conn)
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
