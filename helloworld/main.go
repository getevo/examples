package main

import (
	"github.com/getevo/evo"
	"github.com/gofiber/fiber"
)

type MyModel struct {
	evo.Model
	Name     string
	Username string
}

func main() {
	evo.Setup()

	db := evo.GetDBO()
	db.AutoMigrate(MyModel{})

	evo.Get("test", func(ctx *fiber.Ctx) {
		request := evo.Upgrade(ctx)

		obj := MyModel{
			Name: "Allan", Username: "allan@ies",
		}

		db.Create(&obj)

		request.WriteResponse(obj)
	})

	evo.Get("/", func(ctx *fiber.Ctx) {
		ctx.Write("Hello World")
	})
	evo.Run()
}
