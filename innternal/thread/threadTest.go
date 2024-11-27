package thread

import (
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "8888")
	if err != nil {
		log.Panicln("listen error:", err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println("accept error:", err)
			break
		}
		go Handle(conn)
	}

}
func Handle(conn net.Conn) {
	defer conn.Close()
	packet := make([]byte, 1024)
	for {
		n, err := conn.Read(packet)
		if err != nil {
			log.Println("read socket error:", err)
		}
		_, _ = conn.Write(packet[:n])
	}
}
