package main

import (
	//"conv"
	"fmt"
	"net"
	"stream"
	"time"
)

func main() {
	var errMk error

	st := stream.Reader([]byte("I have a dream"))

	st.WriteBytes([]byte(time.Now().String()))
	st, errMk = st.MakePacket(1)
	if errMk != nil {
		panic(errMk)
	}

	fmt.Println(st.Data())

	con, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	n, err1 := con.Write(st.Data())
	if err1 != nil {
		panic(err1)
	}

	<-time.After(2 * time.Second)

	n, err1 = con.Write(st.Data())
	if err1 != nil {
		panic(err1)
	}
	fmt.Println("size:", n)
	<-time.After(10 * time.Second)
}
