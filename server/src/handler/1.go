package handler

import (
	"fmt"
)

//type PrintMessage struct {
//}

//func (pm *PrintMessage) HandlerOne(send chan []byte, buf []byte) (bool, error) {
//	fmt.Println(string(buf))
//	send <- buf
//	return true, nil
//}

func PrintMessage(send chan []byte, buf []byte) (bool, error) {
	fmt.Println("rec string:", string(buf))
	send <- buf
	return true, nil
}

func init() {
	RegisterHandler(1, PrintMessage)
}
