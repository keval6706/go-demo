package middleware

import (
	"fmt"
	"net/http"

	"github.com/kataras/iris/v12"
	"go.demo/utils"
)

func AuthMiddleware(ctx iris.Context) {
	fmt.Println("Auth Middleware Called")

	// authorization := ctx.GetHeader("Authorization")
	// fmt.Println("Authorization", authorization)

	var hs utils.AuthHeaders
	if err := ctx.ReadHeaders(&hs); err != nil {
		ctx.JSON(iris.Map{
			"error":  err.Error(),
			"status": http.StatusUnauthorized,
		})
		return
	}

	ctx.Values().Set("user", hs)

	ctx.Next()
}
