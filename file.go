package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

type Server struct {
	name string
}

func (s Server) Listen() {
	serv, err := net.Listen("tcp", ":80")
	if err == nil {
		for {
			con, e := serv.Accept()
			p := make([]byte, 1024)
			con.Read(p)
			fmt.Print(string(p))
			ar := strings.Split(string(p), "\n")
			fmt.Print(ar[0])
			if e == nil {
				f, e := os.Open("index.html")
				if e != nil {
					fmt.Print("Error")
				} else {
					text := make([]byte, 1024)
					f.Read(text)
					fmt.Fprintf(con, "HTTP/1.1  200 Ok\nContent-Type:text/html\n\n"+string(text))
				}
			}
			con.Close()
		}
	}
}

func main() {
	s := Server{name: "Juan"}
	s.Listen()
}
