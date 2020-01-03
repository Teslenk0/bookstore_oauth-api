package http

import (
	"github.com/Teslenk0/bookstore_oauth-api/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AccessTokenHandler interface {
	GetById(c *gin.Context)
}

type accessTokenHandler struct {
	service services.Service
}

func NewHandler(s services.Service) AccessTokenHandler {
	return &accessTokenHandler{service: s}
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	access_tokenId := strings.TrimSpace(c.Param("access_token_id"))
	accessToken, err := handler.service.GetById(access_tokenId)
	if err != nil{
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}
