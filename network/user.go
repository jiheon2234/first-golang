package network

import (
	"CRUD-SERVER/service"
	"CRUD-SERVER/types"
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
)

var (
	userRouterInit     sync.Once
	userRouterInstance *userRouter
)

type userRouter struct {
	router *Network

	userService *service.User
}

func newUserRouter(router *Network, userService *service.User) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router:      router,
			userService: userService,
		}

		router.registerGET("/", userRouterInstance.get)
		router.registerPOST("/", userRouterInstance.create)
		router.registerUPDATE("/", userRouterInstance.update)
		router.registerDELETE("/", userRouterInstance.delete)

	})

	return userRouterInstance
}

func (u *userRouter) create(c *gin.Context) {
	fmt.Println("create!!!!!!!!!!!!!!")

	u.userService.Create(nil)

	u.router.okResponse(c, types.NewApiResponse(
		"create Sucesss!!!!", 1))

}
func (u *userRouter) get(c *gin.Context) {
	fmt.Println("get!!!!!!!!!!!!!!")

	u.router.okResponse(c, &types.UserResponse{
		ApiResponse: &types.ApiResponse{Result: 1,
			Description: "get Sucesss!!!!",
		},
		Users: u.userService.Get(),
	})
}
func (u *userRouter) update(c *gin.Context) {
	u.userService.Update(nil, nil)
	fmt.Println("update!!!!!!!!!!!!!!")
}
func (u *userRouter) delete(c *gin.Context) {
	fmt.Println("delete!!!!!!!!!!!!!!")

	u.userService.Delete(nil)
}
