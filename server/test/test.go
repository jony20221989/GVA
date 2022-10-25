package main

import (
	"fmt"
	"os"
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
	fileName := "D:\\BtSoft\\mysql\\MySQL5.6"
	dir, err := os.ReadDir(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, info := range dir {
		fmt.Println(info.Name())
	}

}
