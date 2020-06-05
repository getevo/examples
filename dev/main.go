package main

import (
	"github.com/getevo/evo"
	"github.com/getevo/evo/apps/admin"
	"github.com/getevo/evo/apps/adminlte"
	"github.com/getevo/evo/apps/auth"
	"github.com/getevo/evo/apps/query"
	"github.com/getevo/evo/apps/test"
)

func main() {

	evo.Setup()
	adminlte.Register()
	admin.Register()
	auth.Register()
	query.Register()
	test.Register()
	evo.Run()
}
