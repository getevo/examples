package strange

import (
	"fmt"
	"github.com/CloudyKit/jet"
	"github.com/getevo/evo"
	"github.com/getevo/evo/menu"
	"github.com/getevo/evo/user"
	"github.com/gofiber/fiber"
)

var Path string
var views *jet.Set

// App settings app struct
type App struct{}

// Register register the auth in io apps
func Register() {
	evo.Register(App{})

}

// Register settings app
func (App) Register() {
	fmt.Println("Strange Registered")
	Path = evo.GuessAsset(App{})
	views = evo.RegisterView("strange", Path+"/views")
}

// Router setup routers
func (App) Router() {

	evo.Get("admin/strange", func(ctx *fiber.Ctx) {
		r := evo.Upgrade(ctx)
		r.RenderView(nil, "strange.test")
	})
}

// Permissions setup permissions of app
func (App) Permissions() []user.Permission {
	return []user.Permission{}
}

// Menus setup menus
func (App) Menus() []menu.Menu {
	return []menu.Menu{}
}

// WhenReady called after setup all apps
func (App) WhenReady() {}

func (App) Pack() {
	evo.Pack(Path)
}
