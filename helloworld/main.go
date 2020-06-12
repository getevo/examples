package main

import (
	"github.com/getevo/evo"
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

	evo.Get("test", func(request *evo.Request) {
		obj := MyModel{
			Name: "Allan", Username: "allan@ies",
		}

		db.Create(&obj)

		request.WriteResponse(obj)
	})

	evo.Get("/", func(request *evo.Request) {
		request.Write("Hello World")
	})
	evo.Run()
}
