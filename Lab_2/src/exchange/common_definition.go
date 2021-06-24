package exchange

import (
	"Lab_2/src/cfg"
	"net"
)

type NetWorker struct {
	CliCfg* cfg.ClientConfig
	Comm NetworkCommunication
}

type NetworkCommunication interface {
	DialServer(cfg *cfg.ClientConfig) net.Conn
	CloseConnection(conn net.Conn)

	SendData(data []byte, conn net.Conn) bool
	ReceiveData(conn net.Conn) []byte
}

type DefaultNetworkComm struct {}

type ClientRequest struct {
	ErrType string
	IsCritical bool
}

type ServerRespond struct {
	IsSuccess bool
}

type netData interface {
	ParseRequest(data []byte) (ClientRequest, bool)
	SerializeRequest(request ClientRequest) []byte

	ParseRespond(data []byte) (ServerRespond, bool)
	SerializeRespond(respond ServerRespond) []byte
}