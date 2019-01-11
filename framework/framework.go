package main

import "syscall/js"

func RegisterFunction(funcName string, myfunc func(i []js.Value)) {
	js.Global().Set(funcName, js.NewCallback(myfunc))
}

func Start() {
	println(" Framework Initialized")
}
