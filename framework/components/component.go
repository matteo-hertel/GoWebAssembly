package components

import "syscall/js"

type Component interface {
	Render() string
}

func RegisterFunction(funcName string, myfunc func(i []js.Value)) {

	js.Global().Set(funcName, js.NewCallback(myfunc))
}
