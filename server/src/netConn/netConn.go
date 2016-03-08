package netConn

import (
	//	"bufio"
	"encoding/binary"
	"fmt"
	"handler"
	"io"
	"log"
	"net"
	//	"stream"
	"time"
)

//处理消息的方法

type NetConn struct {
	session    net.Conn
	_exit      chan bool //退出接受消息
	m_rec      chan []byte
	m_send     chan []byte
	disconnect bool
}

const (
	reci_max          = 1000
	send_max          = 1000
	msg_max_length    = 1024 * 1024
	msg_header_length = 2
	connect_timeout   = 5 * time.Minute
)

func (conn *NetConn) WaitForExit() {
	<-conn._exit
	log.Println("exit <- close ")
}

func (c *NetConn) close() {
	c.disconnect = true
	c.session.Close()
}

func (c *NetConn) keepAlive() {
	c.session.SetDeadline(time.Now().Add(connect_timeout)) //设置超时时间
}

func (c *NetConn) read() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Reader Data", err)
		}
		fmt.Println("Reader Data")
		c.close()
		close(c.m_rec)
		close(c.m_send)
		c._exit <- true
		close(c._exit)
	}()

	for {
		var uLenth uint32
		err := binary.Read(c.session, binary.BigEndian, &uLenth)
		if err != nil {
			if err != io.EOF {
				log.Println("Reader Data1:", err)
				panic(err)
			}
			break
		}
		fmt.Println("long :", uLenth)
		if uLenth > msg_max_length {
			log.Printf("Msg length(%v) out of range. \n", uLenth)
			break
		}

		buf := make([]byte, uLenth)
		err = binary.Read(c.session, binary.BigEndian, buf)
		if err != nil {
			if err != io.EOF {
				log.Println("Reader Data2:", err)
			} else {
				panic(err)
			}
		}

		c.keepAlive()
		c.m_rec <- buf
	}
}

func NewNetConn(conn net.Conn) {
	defer func() { conn.Close() }()
	c := new(NetConn)
	c.session = conn
	c._exit = make(chan bool)
	c.m_rec = make(chan []byte, reci_max)
	c.m_send = make(chan []byte, send_max)
	go c.read()
	go c.handlerConn()
	c.WaitForExit()
}

func init() {

}

func (c *NetConn) handlerConn() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("makeMsg : ", err)
		}
		c.close()
	}()

	for {
		buf, ok := <-c.m_rec
		if !ok {
			break
		}
		if len(buf) < msg_header_length {
			break
		}
		handler.HandlerTheMessage(c.m_send, buf)
		fmt.Println(buf)
		//c._exit <- true
	}
}

func (c *NetConn) send() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("sendMessage : ", err)
		}
		c.close()
	}()

	for {
		buf, ok := <-c.m_send
		if !ok {
			break
		}
		if len(buf) < msg_header_length {
			break
		}

		c.session.Write(buf)

		fmt.Println(buf)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
