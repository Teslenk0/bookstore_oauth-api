package app

import (
	"github.com/Teslenk0/bookstore_oauth-api/src/http"
	"github.com/gin-gonic/gin"
)

func mapUrl (router *gin.Engine, atHandler http.AccessTokenHandler){
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.PATCH("/oauth/access_token", atHandler.UpdateExpirationTime)
}