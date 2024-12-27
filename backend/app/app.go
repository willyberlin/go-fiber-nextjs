package app

import (
	"fiber-rooms/app/core"
	"fiber-rooms/app/cors"
	"fiber-rooms/app/rooms"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/graphql-go/handler"
)

func SetupApp(app *fiber.App) {
	cors.SetupCors(app)
	core.SetupRoutes(app)
	rooms.SetupRoutes(app)

	graphqlHandler := handler.New(&handler.Config{
		Schema:   &rooms.Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	fiberGraphQLHandler := adaptor.HTTPHandler(graphqlHandler)
	app.All("/graphql", fiberGraphQLHandler)
}
