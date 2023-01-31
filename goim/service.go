package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Service struct {
	Ip   string
	Port int

	OnlineMap map[string]*User
	mapLock   sync.Mutex

	Message chan string
}

func NewService(ip string, port int) *Service {

	server := &Service{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}

	return server
}

func (s *Service) Handle(conn net.Conn) {
	println("链接建立成功")
	defer conn.Close()
	user := NewUser(s, conn)

	user.Online()

	isLive := make(chan bool)

	go func() {
		buf := make([]byte, 1024)

		for {
			n, err := conn.Read(buf)

			if n == 0 {
				user.Offline()
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err:", err)
				return
			}

			msg := string(buf[:n-1])

			user.DoMsg(msg)

			isLive <- true
		}
	}()

	for {
		select {
		case <-isLive:
		case <-time.After(time.Second * 10):
			user.SendMsg("timed out")
			close(user.ch)
			//关闭连接
			//	time.Sleep(time.Second * 1)

			conn.Close()
			return
		}
	}

}

func (s *Service) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg

	s.Message <- sendMsg
}

func (s *Service) listenMsg() {
	for {
		sendMsg := <-s.Message

		s.mapLock.Lock()

		for _, user := range s.OnlineMap {
			user.ch <- sendMsg
		}
		s.mapLock.Unlock()
	}
}

func (s *Service) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))

	if err != nil {
		fmt.Println("net listen err:", err)
		return
	}
	// Close listener
	defer listener.Close()

	go s.listenMsg()

	for {
		accept, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err : ", err)
			return
		}

		go s.Handle(accept)

	}

}
