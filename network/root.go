package network

import (
	"CRUD-SERVER/service"
	"github.com/gin-gonic/gin"
)

// 라우터설정

type Network struct {
	engin *gin.Engine

	service *service.Service
}

func NewNetwork(service *service.Service) *Network {
	r := &Network{
		engin: gin.New(),
	}

	newUserRouter(r, service.User)

	return r
}

func (r *Network) ServerStart(port string) error {
	return r.engin.Run(port)
}
