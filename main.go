package main

import (
	databases "github.com/Zyprush18/Kasir_Go/src/Databases"
	routes "github.com/Zyprush18/Kasir_Go/src/Routes"
)

func main() {
	databases.Connect()
	routes.Route()
}
