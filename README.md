# go-webassembly
Go WebAssumbly simple demo.
## usage
copy `wasm_exec.js` to current
```
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```
build webassembly binary code(lib.wasm)
```
GOOS=js GOARCH=wasm go build -o lib.wasm main.go
```
## run
listening on :8080
```
go run server.go
```
## reference
- https://github.com/golang/go/wiki/WebAssembly
- https://tutorialedge.net/golang/go-webassembly-tutorial/
