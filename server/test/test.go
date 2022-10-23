package main

import (
	"context"
	"fmt"
	"server/utils"
)

type Result string

func find(ctx context.Context, query string) (Result, error) {
	return Result(fmt.Sprintf("result for %q", query)), nil
}

func main() {
	j1 := utils.NewJWT()
	j2 := utils.NewJWT()

	fmt.Println(j1)
	fmt.Println(j2)
}
