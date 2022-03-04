package main

import (
	"flag"
	"fmt"
)

func main49() {
	host := flag.String("host", "localhost", "Put your database host")
	port := flag.Int("port", 3306, "Put your database port")
	user := flag.String("user", "root", "Put your database user")
	password := flag.String("password", "rahasia", "Put your database password")
	database := flag.String("database", "MySql", "Put your database name")

	flag.Parse()

	fmt.Println("host:", *host)
	fmt.Println("port:", *port)
	fmt.Println("user:", *user)
	fmt.Println("password:", *password)
	fmt.Println("database:", *database)
}
