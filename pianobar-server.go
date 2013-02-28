package main

import (
    "net"
    "fmt"
    "os"
    "io"
    "strings"
)

const listenPort = ":8181"
const ok = "OK"
const valcmd = "npsq"

func main() {
    // connect to the listening port
    fmt.Printf("//+// pianobar control server spinning up //+//\n")

    ln, err := net.Listen("tcp", listenPort)
    defer ln.Close()

    if err != nil {
        fmt.Printf("could not connect to port, err: %v\n", err)
        return
    }

    for {
        // accept the connection
        conn, err := ln.Accept()
        if err != nil {
            fmt.Printf("could not accept connection, err: %v\n", err)
            return
        }

        go func() {
            var req string
            var buf [24]byte

            // read from the connection
            //for {
                n, err := conn.Read(buf[0:])
                if n > 0 {
                    req += string(buf[0:n])
                }
                if err == io.EOF {
                    //break
                }
            //}

            fmt.Printf("i was sent %v\n", req)

            cmd := string(req[0])

            if !strings.Contains(valcmd, cmd) {
                conn.Write([]byte("newp!"))
                return
            }

            if cmd == "s" {
                req = req + "\n"
            }

            // open the fifo file
            fname := "/home/msw978/.config/pianobar/ctl"
            //t := "/home/msw978/pbartmp"
            fi, err := os.OpenFile(fname, os.O_CREATE | os.O_RDWR, 0644)
            if err != nil {
                fmt.Printf("could not open file, err: %v\n", err)
                return
            }

            _, err = fi.Write([]byte(req))
            if err != nil {
                fmt.Printf("could not write to file, err: %v\n", err)
                fi.Close()
                return
            }

            fi.Close()

            conn.Write([]byte(ok))
        }()
    }

}
