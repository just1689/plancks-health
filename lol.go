package main

import (
	"fmt"

	"github.com/docker/docker/client"
)

func Client() {
	cli, err := client.NewEnvClient()
	if err != nil {
		fmt.Println("error")

	}
	cli.Close()
}
