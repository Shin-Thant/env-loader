package main

import (
	"fmt"

	envloader "github.com/Shin-Thant/env-loader"
)

type stc struct {
	PORT         interface{}
	DATABASE_URL string
}

func main() {
	stc := stc{}
	envloader.LoadEnv(&stc)

	fmt.Println(stc)
}
