# Go Web Assembly

Playground for WebAssembly in Go

## Run the program

to target web assembly while compiling the Go code run 

```go
GOARCH=wasm GOOS=js go build -o ./src/lib.wasm main.go
```
