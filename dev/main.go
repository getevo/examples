package main

import (
	"github.com/getevo/evo"
	"github.com/getevo/evo/apps/admin"
	"github.com/getevo/evo/apps/adminlte"
	"github.com/getevo/evo/apps/auth"
	"github.com/getevo/evo/apps/bible"
	"github.com/getevo/evo/apps/query"
	"github.com/getevo/evo/apps/test"
	//"github.com/getevo/examples/strange"
	//"github.com/gofiber/pprof"
)

func main() {

	evo.Setup()
	adminlte.Register()
	//strange.Register()
	admin.Register()
	auth.Register()
	query.Register()
	test.Register()
	bible.Register()
	evo.GetFiber()//.Use(pprof.New())
	evo.Run()
}
