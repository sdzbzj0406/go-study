package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int

	// 在线用户列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	// 消息广播的channel
	Message chan string
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}

	return server
}

// ListenMessage 监听message广播消息channel的goroutine，一旦有消息，就发送给全部的user
func (s *Server) ListenMessage() {

	for true {
		msg := <-s.Message

		// 发送给全部用户
		s.mapLock.Lock()
		for _, cli := range s.OnlineMap {
			cli.C <- msg
		}
		s.mapLock.Unlock()
	}
}

func (s *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	s.Message <- sendMsg
}

func (s *Server) Handler(conn net.Conn) {
	// 当前业务
	fmt.Println("链接建立成功")
	user := NewUser(conn, s)
	//// 用户上线，加入到在线表中
	//s.mapLock.Lock()
	//s.OnlineMap[user.Name] = user
	//s.mapLock.Unlock()
	//
	//// 广播当前用户的上线消息
	//s.BroadCast(user, " is online")

	user.Online()

	isLive := make(chan bool)

	// 接收客户端消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)

			if n == 0 {
				//s.BroadCast(user, "is downing")
				user.offline()
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("conn read error")
				return
			}

			// 提取消息
			msg := string(buf[:n-1])

			// 广播消息
			user.DoMessage(msg)
			//s.BroadCast(user, msg)

			// 用户任意消息，代表是活跃的
			isLive <- true
		}
	}()

	// 当前handle阻塞
	for true {
		select {
		case <-isLive:
			// 重置定时器，不做任何事情，为了激活下面的定时器
		case <-time.After(time.Second * 10):
			// 说明已经超时，将当前客户端强制下线
			user.sendMsg("you are down by timeAfter")
			close(user.C)
			err := conn.Close()
			if err != nil {
				return
			}
			// 退出当前的handler
			return
		}
	}
}

// Start 启动服务器
func (s *Server) Start() {

	// socket listen
	listener, e := net.Listen("tcp4", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if e != nil {
		fmt.Println("net listen error: ", e)
		return
	}

	// close socket
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {

		}
	}(listener)

	go s.ListenMessage()

	for true {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}

		// do handler
		go s.Handler(conn)
	}
}
