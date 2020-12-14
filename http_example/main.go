package main

///
import (
	"github.com/getevo/evo"
	"github.com/getevo/evo/apps/rdb"
)

//
func main() {

	evo.Setup()
	rdb.Register()
	evo.Get("/", func(r *evo.Request) {
		r.WriteResponse("data allan")
	})
	/*func(r *evo.Request) {
		fmt.Println("http Registered")
		r.WriteResponse("Home")
	})*/
	evo.Run()
}
