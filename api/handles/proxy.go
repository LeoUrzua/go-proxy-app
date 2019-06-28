package handles

import (
	"github.com/LeoUrzua/proxy-app/api/middleware"
	"github.com/kataras/iris"
)

// HandlerRedirection should redirect traffic
func HandlerRedirection(app *iris.Application){
	app.Get("/", middleware.Handler, pingHandler)
}

func pingHandler(c iris.Context){
	c.JSON(iris.Map{"Result": "ok"})
}
