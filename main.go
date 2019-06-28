package main

import (
	handlers "github.com/LeoUrzua/proxy-app/api/handles"
	server "github.com/LeoUrzua/proxy-app/api/server"
utils "github.com/LeoUrzua/proxy-app/api/utils"
)

func main() {
	utils.LoadEnv()
	app := server.SetUp()
	handlers.HandlerRedirection(app)
	server.RunServer(app)
}
