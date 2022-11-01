package main

import "fmt"

type Foo struct {
	Data interface{}
}

func (foo Foo) GetData() interface{} {
	return foo.Data
}

func (foo *Foo) SetData(data interface{}) {
	foo.Data = data
}

func main() {
	f := &Foo{}
	f.SetData([]string{"a", "b", "c"})

	var data []string = f.GetData().([]string)
	fmt.Println(data)
}
