package main

import (
	"log"
	"net/http"

	database "go.demo/db"
	"go.demo/routes"
	"google.golang.org/api/iterator"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type PingResponse struct {
	Message string                   `json:"message"`
	Items   []map[string]interface{} `json:"items"`
}

func main() {
	app := iris.New()

	// app.Logger().SetLevel("debug")

	client, fbCtx := database.Firebase()
	defer client.Close()

	// db := database.InitializeDB()
	ent, entCtx := database.Ent()
	app.RegisterDependency(ent)
	defer ent.Close()

	redis := database.Redis()
	app.RegisterDependency(redis)
	defer redis.Close()

	es, _ := database.Elasticsearch()
	app.RegisterDependency(es)

	// Auth Routes
	authRoutes := app.Party("auth")
	mvc.Configure(authRoutes, routes.AuthRouter)

	// Member Routes
	memberRoutes := app.Party("members")
	// memberRoutes.Use(middleware.AuthMiddleware)
	mvc.Configure(memberRoutes, routes.MemberRouter)

	app.Get("ping", func(ctx iris.Context) {
		var items []map[string]interface{}
		documents := client.Collection("attachments").Limit(10).Documents(fbCtx)
		for {
			doc, err := documents.Next()
			if err == iterator.Done {
				break
			}
			if err == nil {
				items = append(items, doc.Data())
			}
		}

		res := PingResponse{
			Message: "pong",
			Items:   items,
		}
		ctx.JSON(res)
	})

	app.Get("db", func(ctx iris.Context) {

		// var members []models.Member
		// db.Debug().Model(&models.Member{}).Find(&members)

		members, err := ent.Debug().Member.Query().Limit(10).All(entCtx)
		if err != nil {
			log.Fatal(err)
		}

		ctx.JSON(iris.Map{
			"status": http.StatusOK,
			"items":  members,
		})
	})

	// Listens and serves incoming http requests
	// on http://localhost:8080.
	app.Listen(":8080")
}

func DemoMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}
