package main

import (
	config "altastore/config"
	routes "altastore/routes"
)

func main() {
	config.InitDb()

	e := routes.New()

	e.Logger.Fatal(e.Start(":8000"))
}
