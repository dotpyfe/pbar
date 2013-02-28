package main

import (
    "fmt"
    "net"
    "os"
    "io"
    "errors"
)

const server = ""

var stations = map[int]string{
    1 : "Albert King",
    2 : "Bonobo",
    3 : "Burial",
    4 : "Eric Clapton",
    5 : "Jimi Hendrix",
    6 : "Phaeleh",
    7 : "The Black Keys",
    8 : "Washed Out"}

func switch_station() (string, error) {
    fmt.Printf("=== CHOOSE STATION ======\n")
    for i, r := range stations {
        fmt.Println(i, " : ", r)
    }
    fmt.Printf("=========================\n")

    var i int
    fmt.Scan(&i)

    if (i > len(stations) - 1) || (i <= 0) {
        fmt.Printf("NEWP\n")
        return "", errors.New("NEWP")
    }

    return stations[i] + " Radio", nil
}

func main() {
    argc := len(os.Args)
    if argc != 2 {
        fmt.Printf("usage: pbarc command\n")
        return
    }

    cmd := os.Args[1]

    var req string
    var err error

    switch cmd {
    case "next", "n":
        req = "n"
    case "pause", "p":
        req = "p"
    case "stop", "q":
        req = "q"
    case "switch", "s":
        req, err = switch_station()
        fmt.Println("chose", req)
        if err != nil {
            return
        }
    default:
        fmt.Println("not a valid command asshole!")
        return
    }

    conn, err := net.Dial("tcp", server)
    if err != nil { panic(err) }

    // TODO: sanity checking
    conn.Write([]byte(req))

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
