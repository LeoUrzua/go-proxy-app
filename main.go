//package proxy_app
package main

import (
	"github.com/LeoUrzua/proxy-app/api/handles"
	"github.com/LeoUrzua/proxy-app/api/middleware"
	"github.com/LeoUrzua/proxy-app/api/server"
	"github.com/LeoUrzua/proxy-app/api/utils"
)

/*
Router Iris
Env vars
 */

func main(){
	utils.LoadEnv()
	app := server.SetUp()
	middleware.InitQueue()
	handles.HandlerRedirection(app)
	server.RunServer(app)
}
