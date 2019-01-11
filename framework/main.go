package main

import (
	"./components"
	"./router"
)

func main() {
	// Starts the Oak framework
	println(" Framework Initialized")

	// Starts our Router
	router.NewRouter()
	router.RegisterRoute("about", components.About)
	// keeps our app running
	done := make(chan struct{}, 0)
	<-done
}
