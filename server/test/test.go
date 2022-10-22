package main

import (
	"context"
	"fmt"
	"time"
)

type Result string

func find(ctx context.Context, query string) (Result, error) {
	return Result(fmt.Sprintf("result for %q", query)), nil
}

func main() {
	var a = 86400000000000

	fmt.Println(int64(a))
	//time.Now().Add(a).Unix()

	fmt.Println(3600 * 24 * time.Duration(1) * time.Second)
}
