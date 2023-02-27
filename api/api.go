package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mgmt-server/dal"
)

type API struct {
	router      *gin.Engine
	userDAL     *dal.UserDAL
	chatRoomDAL *dal.ChatRoomDAL
}

func NewAPI(userDAL *dal.UserDAL, chatRoomDAL *dal.ChatRoomDAL) *API {
	router := gin.Default()
	api := &API{
		router:      router,
		userDAL:     userDAL,
		chatRoomDAL: chatRoomDAL,
	}
	api.registerRoutes()
	return api
}

func (api *API) Start(port int) error {
	return api.router.Run(fmt.Sprintf(":%d", port))
}

func (api *API) registerRoutes() {
	v1 := api.router.Group("/v1")
	v1.POST("/users", api.createUser)
	v1.POST("/chatrooms", api.createChatRoom)
}
