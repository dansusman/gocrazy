package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

    host := "127.0.0.1"
    port := "33333"
    address := fmt.Sprintf("%s:%s", host, port)

    conn, err := makeConn(address)

    fatalError(err)

    defer conn.Close()

    fmt.Print("Text to send: ")
    reader := bufio.NewReader(os.Stdin)

    text, readErr := reader.ReadString('\n')
    fatalError(readErr)

    _, writeErr := conn.Write([]byte(text))

    conReader := bufio.NewReader(conn)
    res, resErr := conReader.ReadString('\n')
    fatalError(resErr)

    fmt.Println("Server says: " + res)

    fatalError(writeErr)

}

func makeConn(address string) (net.Conn, error) {
    return net.Dial("tcp", address)
}

func fatalError(err error) {
    if err != nil {
        panic("Client error: " + err.Error())
    }
}
