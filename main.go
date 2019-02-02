package main

import (
  "strconv"
  "syscall/js"
)

func add(ids []js.Value) {
  value1 := js.Global().Get("document").Call("getElementById", ids[0].String()).Get("value").String()
  value2 := js.Global().Get("document").Call("getElementById", ids[1].String()).Get("value").String()

  int1, _ := strconv.Atoi(value1)
  int2, _ := strconv.Atoi(value2)

  js.Global().Get("document").Call("getElementById", ids[2].String()).Set("value", int1+int2)
}

func registerCallbacks() {
  js.Global().Set("add", js.NewCallback(add))
}

func main() {
  c := make(chan struct{}, 0)
  println("Go WebAssembly Initialized!")
  registerCallbacks()

  <-c
}
