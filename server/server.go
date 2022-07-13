package main

import (
    "fmt"
    "net"
)

func main() {

    host := "127.0.0.1"
    port := "33333"

    address := fmt.Sprintf("%s:%s", host, port)

    l, err := net.Listen("tcp", address)

    fatalError(err)

    defer l.Close()

    fmt.Printf("Listening on %s\n", address)

    for {
        // we're spinnin' forever, listening for stuff
        // TODO: let's add a clock next, and send a "BYE" message
        // after 10 seconds
        // client will run until "BYE" is received!

        conn, err := l.Accept()

        fatalError(err)

        go func(conn net.Conn) {
            buf := make([]byte, 1024)
            len, err := conn.Read(buf)

            if err != nil {
                fmt.Printf("Error reading: %#v\n", err)
                return
            }

            fmt.Printf("Message received: %s\n", string(buf[:len]))

            conn.Write([]byte("Message received. Thanks!\n"))
            conn.Close()
        }(conn)
    }
}


func fatalError(err error) {
    if err != nil {
        panic("Server err: " + err.Error())
    }
}
