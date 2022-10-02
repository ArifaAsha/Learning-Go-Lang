package main
import "fmt"

type Fooer interface {
	Foo() string
  }
  
  type Container struct {
	Fooer
  }