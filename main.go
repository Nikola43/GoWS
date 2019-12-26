package main

import "github.com/nikola43/testsocket/controllers"

func main() {
	a := controllers.App{}
	a.Initialize()
	a.Run(":8080")
}
