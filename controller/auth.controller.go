package controller

import (
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"go.demo/utils"
)

type AuthController struct {
}

func (m *AuthController) BeforeActivation(b mvc.BeforeActivation) {

}

// Get - members/profile/followers
func (m *AuthController) GetToken(ctx iris.Context) {
	claims := utils.TokenClaims{Id: uuid.New().String(), Name: "Keval G"}

	token, _ := utils.GenerateToken(claims)
	data, _ := utils.VerifyToken(string(token))

	ctx.JSON(iris.Map{"token": string(token), "data": data})
}
