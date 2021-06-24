package cfg

import (
	toml "github.com/pelletier/go-toml"
)

type ServerCredentials struct {
	Ip string
	Port int
}

type ClientConfig struct {
	CriticalErrorsList []string
	SrvCred ServerCredentials
}

type ServerConfig struct {
	Port int
}

func ParseSrvCfg(cfgPath string) (ServerConfig, bool) {
	tree, err := toml.LoadFile(cfgPath)
	if err != nil {
		return ServerConfig{}, false
	}
	srvCfg := ServerConfig{
		Port: int(tree.GetDefault("server.port", 4500).(int64)),
	}
	return srvCfg, true
}

func ParseCliCfg(cfgPath string) (ClientConfig, bool) {
	tree, err := toml.LoadFile(cfgPath)
	if err != nil {
		return ClientConfig{}, false
	}
	var cliCfg ClientConfig
	cliCfg.SrvCred = parseSrvCredTable(tree)
	cliCfg.CriticalErrorsList = parseCriticalErrorsList(tree)
	return cliCfg, true
}

func parseSrvCredTable(tree *toml.Tree) ServerCredentials {
	var srvCred ServerCredentials
	srvCred.Ip = tree.GetDefault("server.ip", "127.0.0.1").(string)
	srvCred.Port = int(tree.GetDefault("server.port", 4500).(int64))
	return srvCred
}

func parseCriticalErrorsList(tree *toml.Tree) []string {
	return tree.GetArray("errors.critical").([]string)
}