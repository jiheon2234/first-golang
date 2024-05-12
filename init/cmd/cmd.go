package cmd

import (
	"CRUD-SERVER/config"
	"CRUD-SERVER/network"
	"CRUD-SERVER/repository"
	"CRUD-SERVER/service"
	"fmt"
)

// 다른 항목들을 모두 받아서 시작환경 구성?

type Cmd struct {
	config     *config.Config
	network    *network.Network
	repository *repository.Repository
	service    *service.Service
}

func NewCmd(filePath string) *Cmd {
	c := &Cmd{
		config: config.NewConfig(filePath),
	}

	c.repository = repository.NewRepository()
	c.service = service.NewService(c.repository)
	c.network = network.NewNetwork(c.service)

	fmt.Println(c.config.Server.Port)

	c.network.ServerStart(c.config.Server.Port)

	return c
}
