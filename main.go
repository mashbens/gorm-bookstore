package main

import (
	"books/conf"
	"books/route"
)

func main() {
	conf.InitDB()
	e := route.New()
	e.Logger.Fatal(e.Start(":8080"))
}
