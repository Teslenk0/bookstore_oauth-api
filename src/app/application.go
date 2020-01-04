package app

import (
	"github.com/Teslenk0/bookstore_oauth-api/src/http"
	"github.com/Teslenk0/bookstore_oauth-api/src/repository/db"
	"github.com/Teslenk0/bookstore_oauth-api/src/repository/rest"
	"github.com/Teslenk0/bookstore_oauth-api/src/services"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {

	atHandler := http.NewAccessTokenHandler(access_token.NewService(rest.NewRestUsersRepository(), db.NewRepository()))

	mapUrl(router, atHandler)

	_ = router.Run(":8080")
}
