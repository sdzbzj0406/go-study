package main

import (
	"fmt"
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn

	server *Server
}

// NewUser 创建用户
func NewUser(conn net.Conn, server *Server) *User {

	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}

	go user.ListenMessage()

	return user
}

// Online 上线
func (u *User) Online() {

	// 用户上线，加入到在线表中
	u.server.mapLock.Lock()
	u.server.OnlineMap[u.Name] = u
	u.server.mapLock.Unlock()

	// 广播当前用户的上线消息
	u.server.BroadCast(u, " is online")
}

// offline 下线
func (u *User) offline() {
	// 用户上线，加入到在线表中
	u.server.mapLock.Lock()
	delete(u.server.OnlineMap, u.Name)
	u.server.mapLock.Unlock()

	// 广播当前用户的上线消息
	u.server.BroadCast(u, " is downing")
}

// 发送消息
func (u *User) sendMsg(msg string) {
	_, err := u.conn.Write([]byte(msg))
	if err != nil {
		return
	}
}

// DoMessage 处理消息
func (u *User) DoMessage(msg string) {

	if msg == "who" {
		// 查询当前在线用户
		u.server.mapLock.Lock()
		for _, user := range u.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + " online\n"
			u.sendMsg(onlineMsg)
		}
		u.server.mapLock.Unlock()

	} else if len(msg) > 7 && msg[:7] == "rename|" {
		// 消息格式： rename|bob
		newName := strings.Split(msg, "|")[1]

		// 判断name是否存在
		_, ok := u.server.OnlineMap[newName]
		if ok {
			u.sendMsg("name has used")
		} else {
			u.server.mapLock.Lock()
			delete(u.server.OnlineMap, u.Name)
			u.server.OnlineMap[newName] = u
			u.server.mapLock.Unlock()

			u.Name = newName
			u.sendMsg("has update name:" + newName + "\n")
		}
	} else if len(msg) > 4 && msg[:3] == "to|" {
		// 私聊信息
		// 1.获取用户名
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			u.sendMsg("msg format error,please check")
			return
		}
		// 2.得到user对象
		remoteUser, ok := u.server.OnlineMap[remoteName]
		if !ok {
			u.sendMsg("user not exist")
			return
		}

		// 3.获取消息内容
		content := strings.Split(msg, "|")[2]
		if content == "" {
			u.sendMsg("no content in msg")
			return
		}

		// 4.通过user对象发送消息
		remoteUser.sendMsg(u.Name + " say : " + content)

	} else {
		u.server.BroadCast(u, msg)
	}
}

// ListenMessage  监听user channel 消息
func (u *User) ListenMessage() {
	for true {
		msg := <-u.C
		write, err := u.conn.Write([]byte(msg + "\n"))
		if err != nil {
			fmt.Println("write error")
			return
		} else {
			fmt.Println(write)
		}
	}
}
