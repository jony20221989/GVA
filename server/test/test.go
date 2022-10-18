package main

import (
	"context"
	"fmt"
)

type Result string

func find(ctx context.Context, query string) (Result, error) {
	return Result(fmt.Sprintf("result for %q", query)), nil
}

func main() {
	//var g singleflight.Group
	//const n = 5
	//waited := int32(n)
	//done := make(chan struct{})
	//key := "https://weibo.com/1227368500/H3GIgngon"
	//for i := 0; i < n; i++ {
	//	go func(j int) {
	//		v, _, shared := g.Do(key, func() (interface{}, error) {
	//			ret, err := find(context.Background(), key)
	//			return ret, err
	//		})
	//		if atomic.AddInt32(&waited, -1) == 0 {
	//			close(done)
	//		}
	//		fmt.Printf("index: %d, val: %v, shared: %v\n", j, v, shared)
	//	}(i)
	//}
	//
	//select {
	//case <-done:
	//case <-time.After(time.Second):
	//	fmt.Println("Do hangs")
	//}
	//定义一个正常变量
	var x int = 5748

	//指针声明
	var p *int

	//初始化指针
	p = &x

	//显示结果
	fmt.Println("存储在x中的值 = ", x)
	fmt.Println("x的内存地址 = ", &x)
	fmt.Println("存储在变量p中的值 = ", p)
	fmt.Println("存储在变量p中的值 = ", *p)
}
