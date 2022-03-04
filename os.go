package main

import (
	"fmt"
	"os"
)

func main48() {
	args := os.Args

	fmt.Println(args)

	hostname, err := os.Hostname()

	if err != nil {
		fmt.Println("error:", err.Error())
	} else {
		fmt.Println("hostname:", hostname)
	}
}
