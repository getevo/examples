package main

///
import (
	"fmt"
	"github.com/getevo/evo"
)

///
func main() {
	evo.Setup()
	evo.Get("/", func(r *evo.Request) {

		fmt.Println("DB Examples Registered")

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
