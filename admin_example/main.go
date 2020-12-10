package main

import (
	"github.com/getevo/evo"
	"github.com/getevo/evo/apps/adminlte"
)

func main() {

	evo.Setup()
	adminlte.Register()

}
