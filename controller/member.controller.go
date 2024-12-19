package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-redis/redis/v8"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"go.demo/utils"
)

type MemberController struct {
}

func (c *MemberController) GetAll() bool {
	return true
}

func (m *MemberController) BeforeActivation(b mvc.BeforeActivation) {
	// b.Dependencies().Add/Remove
	// b.Router().Use/UseGlobal/Done
	// and any standard Router API call you already know
	// 1-> Method
	// 2-> Path
	// 3-> The controller's function name to be parsed as handler
	// 4-> Any handlers that should run before the MyCustomHandler

	// For Register Custom Path or Add Middleware
	// b.Handle("GET", "/something", "MyCustomHandler")

	// For Controller Level Middleware
	// b.Router().Use(func(ctx iris.Context) {
	// 	fmt.Println("Before Activation")
	// 	ctx.Next()
	// })

}

type User struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
}

func (m *MemberController) Post(ctx iris.Context, es *elasticsearch.Client, redis *redis.Client) {
	var user User
	err := ctx.ReadJSON(&user)
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Validation error").
			Detail("One or more fields failed to be validated").
			Type("/user/validation-errors"))

		return
	}

	c := context.Background()
	_u, _ := json.Marshal(user)
	err = redis.HSet(c, "members", user.Email, _u).Err()
	if err != nil {
		errMsg := fmt.Errorf("%v", err)
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Internal Server Error").
			Detail("Failed to insert document into Elasticsearch").
			Type("/internal-server-error").
			DetailErr(errMsg))

		return
	}

	data, _ := json.Marshal(user)
	index, err := es.Index("members",
		bytes.NewReader(data),
		es.Index.WithContext(context.Background()),
		es.Index.WithDocumentID(user.Email),
		es.Index.WithRefresh("true"),
	)
	if err != nil {
		errMsg := fmt.Errorf("indexing document into elastic search: %v", err)
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("Internal Server Error").
			Detail("Failed to insert document into Elasticsearch").
			Type("/internal-server-error").
			DetailErr(errMsg))

		return
	}
	defer index.Body.Close()

	ctx.JSON(iris.Map{"message": "Member added."})
}

func (m *MemberController) GetBy(email string, ctx iris.Context, es *elasticsearch.Client, redis *redis.Client) {

	c := context.Background()
	r, _ := redis.HGet(c, "members", email).Result()

	_member, _ := es.Get("members", email)

	var doc map[string]interface{}
	json.NewDecoder(_member.Body).Decode(&doc)

	ctx.JSON(iris.Map{
		"message": "Member Find By.",
		"r":       r,
		"doc":     doc,
	})
}

func (m *MemberController) MyCustomHandler() string {
	return "MyCustomHandler says Hey"
}

// Get - members/profile/followers
func (m *MemberController) GetProfileFollowers(ctx iris.Context) {
	authorization := ctx.GetHeader("Authorization")
	fmt.Println("Authentication", authorization)

	// var user utils.AuthHeaders
	user, _ := ctx.Values().Get("user").(utils.AuthHeaders)

	ctx.JSON(iris.Map{"user": user})
}
