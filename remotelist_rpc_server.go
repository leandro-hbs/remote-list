package main

import (
	"fmt"
	"net"
	"net/rpc"
	"ppgti/remotelist/pkg"
)

func main() {
	matriz := new(remotelist.RemoteList)
	matriz.LoadMatriz()
	rpcs := rpc.NewServer()
	rpcs.Register(matriz)
	l, e := net.Listen("tcp", "[localhost]:5000")
	defer l.Close()
	if e != nil {
		fmt.Println("listen error:", e)
	}
	for {
		conn, err := l.Accept()
		if err == nil {
			go rpcs.ServeConn(conn)
		} else {
			break
		}
	}
}
