package main

import "fmt"

type Bar interface {
	Get() string
}

func print(i Bar) {
	fmt.Println(i.Get())
}

type Foo struct {
	f string
}

func (foo *Foo) Get() string {
	return foo.f
}

func main() {
	foo := Foo{"hello"}
	print(&foo)
}
