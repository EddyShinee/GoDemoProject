package v2

import (
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Username string `json:"username" json:"username" binding:"required"`
	Password string `json:"password" json:"password" binding:"required"`
}

func GetUsers(c *gin.Context) {
	appG := app.Gin{C: c}
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
	return
}
