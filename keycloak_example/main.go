package main

import (
	"github.com/getevo/evo"
	"github.com/getevo/evo/apps/keycloak"
)

func main() {

	evo.Setup()
	keycloak.Register("https://authfootters.ies-italia.it", "footters", "dev")
	evo.Run()
}
