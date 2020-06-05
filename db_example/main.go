package main

///
import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/iesreza/io"
)

///
func main() {
	evo.Setup()
	evo.Get("/", func(ctx *fiber.Ctx) {

		fmt.Println("DB Examples Registered")

		r := evo.Upgrade(ctx)
		r.Download()
		//b = evo.GetDBO()
		/*type User struct {
			username  	string
			firstname  	string
		}
		user_response = User{
			"allan_nava"
			"Allan"
		}*/
		//r.WriteResponse(user_response)*/
		r.WriteResponse("Trying to add db")
	})
	evo.Run()
}

///
