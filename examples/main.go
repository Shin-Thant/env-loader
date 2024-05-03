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

	//* with custom path
	err := envloader.LoadEnv(&stc, &envloader.LoadEnvOptions{
		EnvPath: "/.env.development",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stc)

	//* with default path (.env)
	err = envloader.LoadEnv(&stc, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stc)
}
