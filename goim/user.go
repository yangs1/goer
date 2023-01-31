package main

import (
	"fmt"
	"net"
)

type User struct {
	Name string
	Addr string
	ch   chan string
	conn net.Conn

	service *Service
}

func NewUser(service *Service, conn net.Conn) *User {
	remoteAddr := conn.RemoteAddr().String()

	user := &User{
		Name:    remoteAddr,
		Addr:    remoteAddr,
		ch:      make(chan string),
		conn:    conn,
		service: service,
	}

	go user.listenConn()

	return user
}

func (user *User) listenConn() {
	for {
		msg := <-user.ch

		write, err := user.conn.Write([]byte(msg + "\n"))
		if err != nil {
			fmt.Println("user write err :", err)
			return
		}
		fmt.Println(fmt.Sprintf("write : %d", write))
	}
}

func (user *User) Online() {
	user.service.mapLock.Lock()

	user.service.OnlineMap[user.Addr] = user

	user.service.BroadCast(user, "已上线")

	user.service.mapLock.Unlock()
}

func (user *User) Offline() {

	user.service.mapLock.Lock()

	delete(user.service.OnlineMap, user.Addr)

	user.service.BroadCast(user, "下线")

	user.service.mapLock.Unlock()

}

func (user *User) DoMsg(msg string) {
	if msg == "who" {
		for _, u := range user.service.OnlineMap {
			sendMsg := "[" + u.Addr + "]" + u.Name + "在线..."
			user.SendMsg(sendMsg)
		}
	} else if len(msg) > 7 && msg[:7] == "rename:" {
		//strings.Split() // explode
		newName := msg[7:len(msg)]

		_, ok := user.service.OnlineMap[newName]
		if ok {
			user.SendMsg("当前用户名称已被使用")
		} else {
			user.service.mapLock.Lock()
			delete(user.service.OnlineMap, user.Name)
			user.Name = newName
			user.service.OnlineMap[newName] = user
			user.service.mapLock.Unlock()
			user.SendMsg("修改名称成功")

		}
	} else {
		user.service.BroadCast(user, msg)
	}
}

func (user *User) SendMsg(msg string) {
	//user.ch <- msg
	user.conn.Write([]byte(msg))
}
