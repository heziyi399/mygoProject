package pkg

import "net"

type ITcpClient interface {
	Connect() error
	// Send 推送消息
	Send() error
	// Message 消息回调
	Message(data []byte) error
	// Close 关闭连接
	Close() error
}
type TcpClient struct {
	address string
	conn    net.Conn
}

func NewTcpClient() *TcpClient {
	return &TcpClient{}
}
