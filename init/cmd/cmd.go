package cmd

import (
	"CRUD-SERVER/config"
	"CRUD-SERVER/network"
	"fmt"
)

// 다른 항목들을 모두 받아서 시작환경 구성?

type Cmd struct {
	config  *config.Config
	network *network.Network
}

func NewCmd(filePath string) *Cmd {
	c := &Cmd{
		config:  config.NewConfig(filePath),
		network: network.NewNetwork(),
	}

	fmt.Println(c.config.Server.Port)

	c.network.ServerStart(c.config.Server.Port)

	return c
}
