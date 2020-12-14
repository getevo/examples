package main

///
import (
	"github.com/getevo/evo"
	"github.com/getevo/evo/apps/rdb"
)

//
func main() {
<<<<<<< HEAD

	evo.Setup()
	rdb.Register()
	evo.Get("/", func(r *evo.Request) {
		r.WriteResponse("data allan")
=======
	evo.Setup()
	rdb.Register()
	evo.Get("/", func(r *evo.Request) {
		r.WriteResponse("data")
>>>>>>> 66d2ae2e9f9e3ab507d4d6be8a98aedc5168418b
	})
	/*func(r *evo.Request) {
		fmt.Println("http Registered")
		r.WriteResponse("Home")
	})*/
	evo.Run()
}
