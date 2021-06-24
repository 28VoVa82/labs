package main

import (
	"Lab_2/src/cfg"
	"Lab_2/src/exchange"
	"fmt"
	"log"
	"net"
)

const configFilePath = "./config/cli_cfg.toml"

func main() {
	srvCfg, res := cfg.ParseSrvCfg(configFilePath)
	if !res {
		log.Printf("Error while parsing client config file %s", configFilePath)
		return
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", srvCfg.Port))
	if err != nil {
		log.Println(err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is running...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			conn.Close()
			return
		}
		handleConnection(conn)
		conn.Close()
	}
}

func handleConnection(conn net.Conn) {
	input := make([]byte, 1024)
	n, err := conn.Read(input)
	if n == 0 || err != nil {
		log.Printf("Reading client failed. Error: %s", err)
		return
	}

	_, isParseOk := exchange.ParseRequest(input[0:n])
	if !isParseOk {
		return
	}

	resp := exchange.SerializeRespond(exchange.ServerRespond{IsSuccess: true})
	if resp == nil {
		return
	}

	_, err = conn.Write(resp)
	if err != nil {
		log.Printf("Server data writing failed. Error: %s", err)
		return
	}
}
