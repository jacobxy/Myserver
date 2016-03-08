package handler

import (
	//"netConn"
	"errors"
	"log"
	"stream"
)

//type Handler interface {
//	HandlerOne(send chan []byte, buf []byte) (bool, error)
//}

type Handler func(send chan []byte, buf []byte) (bool, error)

var Handlers map[uint32]Handler

func GetHandlers() map[uint32]Handler {
	if Handlers == nil {
		Handlers = make(map[uint32]Handler, 1000)
	}
	return Handlers
}

func RegisterHandler(key uint32, hand Handler) bool {
	_, ok := GetHandlers()[key]
	if ok {
		return false
	}
	Handlers[key] = hand
	return true
}

func HandlerTheMessage(send chan []byte, buf []byte) error {
	st := stream.Reader(buf)
	index, err := st.ReadU16()
	if err != nil {
		panic("reader header error")
	}
	r, ok := GetHandlers()[uint32(index)]
	if !ok {
		return errors.New("There is no method ")
	}

	log.Println("method function ")
	r(send, st.DataLeft())
	return nil
}
