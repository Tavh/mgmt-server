package api

import (
	"github.com/gin-gonic/gin"
	"mgmt-server/model"
	"net/http"
)

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

func respondWithError(statusCode int, err error, c *gin.Context) {
	errResponse := &errorResponse{ErrorMessage: err.Error()}
	c.JSON(statusCode, errResponse)
}

func (api *API) createUser(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		respondWithError(http.StatusBadRequest, err, c)
		return
	}

	result, err := api.userDAL.CreateUser(user)
	if err != nil {
		respondWithError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (api *API) createChatRoom(c *gin.Context) {
	var chatRoom model.ChatRoom
	if err := c.BindJSON(&chatRoom); err != nil {
		respondWithError(http.StatusBadRequest, err, c)
		return
	}

	result, err := api.chatRoomDAL.CreateChatRoom(chatRoom)
	if err != nil {
		respondWithError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusCreated, result)
}
