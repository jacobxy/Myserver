package main

import (
	//"bufio"
	//"fmt"
	"net"
	//"error"
	//"execl"
	"netConn"
)

func main() {

	//execl.LoadExecl()
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	checkErr(err)
	for {
		conn, err := listen.Accept()
		checkErr(err)
		netConn.NewNetConn(conn)
	}
}
func checkErr(err error) {
	if err != nil {
		panic("listen error")
	}
}
