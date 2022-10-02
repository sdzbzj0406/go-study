package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int // 当前客户模式
}

func NewClient(serverIp string, serverPort int) *Client {

	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999,
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))

	if err != nil {
		fmt.Println("net dial err:", err)
		return nil
	}

	client.conn = conn

	return client
}

var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "IP address")
	flag.IntVar(&serverPort, "port", 8888, "port address")
}

func (client *Client) menu() bool {

	var f int
	fmt.Println("1.public")
	fmt.Println("2.private")
	fmt.Println("3.rename")
	fmt.Println("0.exit")

	_, err := fmt.Scanln(&f)
	if err != nil {
		return false
	}

	if f >= 0 && f <= 3 {
		client.flag = f
		return true
	} else {
		fmt.Println("menu error")
		return false
	}
}

func (client *Client) UpdateName() bool {
	fmt.Println("please input userName")
	_, err := fmt.Scanln(&client.Name)
	if err != nil {
		return false
	}

	sendMsg := "rename|" + client.Name + "\n"
	_, err = client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn write error")
		return false
	} else {
		return true
	}
}

func (client *Client) DealResponse() {
	// 一旦有消息，就拷贝到标准输出,永久阻塞监听
	_, err := io.Copy(os.Stdout, client.conn)
	if err != nil {
		return
	}

}

func (client *Client) publicChat() {

	var chartMsg string
	// 提示用户发送消息
	fmt.Println("please input msg, exit can end")
	_, err := fmt.Scanln(&chartMsg)
	if err != nil {
		return
	}

	// 发送服务器
	for chartMsg != "exit" {
		if len(chartMsg) != 0 {
			sendMsg := chartMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn write err:", err)
				break
			}
		}
		chartMsg = ""
		fmt.Println("please input msg, exit can end")
		_, err := fmt.Scanln(&chartMsg)
		if err != nil {
			return
		}
	}
}

func (client *Client) SelectUser() {
	sendMsg := "who\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn write error")
	}
}

func (client *Client) privateChat() {

	var remoteName string
	var chartMsg string

	client.SelectUser()
	fmt.Println("please chat to user name")
	_, err := fmt.Scanln(&remoteName)
	if err != nil {
		return
	}

	for remoteName != "exit" {
		// 提示用户发送消息
		fmt.Println("please input msg, exit can end")
		_, err := fmt.Scanln(&chartMsg)

		if err != nil {
			return
		}

		for chartMsg != "exit" {
			if len(chartMsg) != 0 {
				sendMsg := "to|" + remoteName + "|" + chartMsg + "\n\n"

				_, err := client.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("conn write err:", err)
					break
				}
			}

			chartMsg = ""
			fmt.Println("please input msg, exit can end")
			_, err := fmt.Scanln(&chartMsg)
			if err != nil {
				return
			}
		}
		client.SelectUser()
		fmt.Println("please chat to user name")
		_, err = fmt.Scanln(&remoteName)
		if err != nil {
			return
		}
	}
}

func (client *Client) Run() {

	for client.flag != 0 {
		for client.menu() != true {
		}
		// 根据不同模式，处理不同业务
		switch client.flag {
		case 1:
			fmt.Println("public mode")
			client.publicChat()
			break
		case 2:
			fmt.Println("private mode")
			client.privateChat()
			break
		case 3:
			fmt.Println("rename mode")
			client.UpdateName()
			break
		}
	}
	return
}

func main() {

	// 命令行解析
	flag.Parse()

	client := NewClient(serverIp, serverPort)

	if client == nil {
		fmt.Println("client connect server error")
		return
	} else {
		fmt.Println("client connect server success")
	}

	// 启动客户端业务
	//select {}

	// 单独开启，处理server返回的消息
	go client.DealResponse()

	client.Run()
}
