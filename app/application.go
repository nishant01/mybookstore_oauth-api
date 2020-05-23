package app

import (
	"github.com/gin-gonic/gin"
	"github.com/nishant01/mybookstore_oauth-api/clients/cassandra"
	"github.com/nishant01/mybookstore_oauth-api/domain/access_token"
	"github.com/nishant01/mybookstore_oauth-api/http"
	"github.com/nishant01/mybookstore_oauth-api/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session, dBErr := cassandra.GetSession()
	if dBErr != nil {
		panic(dBErr)
	}
	session.Close()
	atHandler := http.NewHandler(access_token.NewService(db.NewRepository()))
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.Run(":8080")
}
