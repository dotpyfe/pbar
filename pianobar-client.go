package main

import (
    "fmt"
    "net"
    "os"
    "io"
)

const server = ""

func main() {
    argc := len(os.Args)
    if argc != 2 {
        fmt.Printf("usage: pbarc command\n")
        return
    }

    cmd := os.Args[1]

    conn, err := net.Dial("tcp", server)
    if err != nil { panic(err) }

    // TODO: sanity checking
    conn.Write([]byte(cmd))

    if err != nil { panic(err) }

    var buf [5]byte
    for {
        n, err := conn.Read(buf[0:])
        if n > 0 {
            fmt.Println("successfully sent", string(buf[0:n]))
            return
        }
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Printf("ermagerd err: %v\n", err)
            return
        }
    }

    conn.Close()
}
