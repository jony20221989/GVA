package main

import (
	"fmt"
	"sync"
)

type Instance struct{}

var (
	once     sync.Once
	instance *Instance
)

func NewInstance() *Instance {
	once.Do(func() {
		instance = &Instance{}
		fmt.Println("instance")

	})
	fmt.Println("Outside")
	return instance
}

func main() {
	//for i := 0; i < 3; i++ {
	//	_ = NewInstance()
	//}

	var a int = 1
	p := &a
	fmt.Printf("%v\n", a)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", p)
	fmt.Printf("%p\n", &p)

}
