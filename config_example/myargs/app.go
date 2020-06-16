package myargs

import (
	"fmt"
	"github.com/getevo/evo"
	"github.com/getevo/evo/lib/log"
	"github.com/getevo/evo/menu"
	"github.com/getevo/evo/user"
)

// Register register the admin in io apps
func Register() {
	fmt.Println("My Args Registered")
	evo.Register(App{})
}

// App admin app struct
type App struct{}

// Custom args
type MyArgs struct {
	MyArg1 string `arg:"-d" help:"A test arg" default:"testarg.yml"`
	MyArg2 bool   `arg:"-e" help:"Copy assets to build dir"`
}

var args = MyArgs{}

// Register register the app
func (App) Register() {
	//set arg to be parsed
	evo.ParseArg(&args)
}

// WhenReady called after setup all apps
func (App) WhenReady() {
	// arg could be accessible here
	log.InfoF("%+v", args)
}

// Router setup routers
func (App) Router() {}

// Permissions setup permissions of app
func (App) Permissions() []user.Permission {
	return []user.Permission{}
}

// Menus setup menus
func (App) Menus() []menu.Menu {
	return []menu.Menu{}
}

func (App) Pack() {}
