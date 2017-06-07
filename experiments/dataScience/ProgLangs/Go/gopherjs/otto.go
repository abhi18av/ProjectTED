package main

import (
  "github.com/robertkrimen/otto"
  "github.com/robertkrimen/otto/repl"
)

func main() {
  vm := otto.New()

  if err := repl.Run(vm); err != nil {
    panic(err)
  }
}
