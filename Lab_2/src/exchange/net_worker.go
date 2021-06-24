package exchange

import (
	"Lab_2/src/cfg"
	"fmt"
	"log"
	"net"
)

func NewNetWorker(config *cfg.ClientConfig, comm NetworkCommunication) *NetWorker {
	return &NetWorker{
		CliCfg:	config,
		Comm:	comm,
	}
}

func (w *NetWorker) SendRecvData(data []byte) []byte {
	conn := w.Comm.DialServer(w.CliCfg)
	defer w.Comm.CloseConnection(conn)

	sendRes := w.Comm.SendData(data, conn)
	if !sendRes {
		return nil
	}

	return w.Comm.ReceiveData(conn)
}

func (d *DefaultNetworkComm) DialServer(cfg *cfg.ClientConfig) net.Conn {
	conn, err := net.Dial("tcp",fmt.Sprintf("%s:%d", cfg.SrvCred.Ip, cfg.SrvCred.Port))
	if err != nil{
		log.Printf("Connection to server failed. Error: %s", err.Error())
		return nil
	}
	return conn
}

func (d *DefaultNetworkComm) CloseConnection(conn net.Conn) {
	conn.Close()
}

func (d *DefaultNetworkComm) SendData(data []byte, conn net.Conn) bool {
	if data == nil{
		log.Printf("No data to send")
		return false
	}

	sendBytes, err := conn.Write(data)
	if err != nil {
		log.Printf("Sending data to server failed. Error: %s", err.Error())
		return false
	}
	if sendBytes == 0 {
		log.Printf("Sending failed")
		return false
	}

	return true
}

func (d *DefaultNetworkComm) ReceiveData(conn net.Conn) []byte {
	buff := make([]byte, 1024)
	recvBytes, err := conn.Read(buff)
	if err != nil {
		log.Printf("Receiving data from server failed. Error: %s", err.Error())
		return nil
	}

	if recvBytes == 0 {
		log.Printf("No bytes recieved from net")
		return nil
	}

	return buff[0:recvBytes]
}