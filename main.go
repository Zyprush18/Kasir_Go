package main

import (
    "github.com/Zyprush18/Kasir_Go.git/src/DataBases"
    "github.com/Zyprush18/Kasir_Go.git/src/Routes"
)

func main() {
    databases.Connect()
    routes.Route()
}