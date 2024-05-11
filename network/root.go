package network

import "github.com/gin-gonic/gin"

// 라우터설정

type Network struct {
	engin *gin.Engine
}

func NewNetwork() *Network {
	r := &Network{
		engin: gin.New(),
	}

	newUserRouter(r)

	return r
}

func (r *Network) ServerStart(port string) error {
	return r.engin.Run(port)
}
