package helper

import "fmt"

var version = "1.0.0"
var Application = "Go"

func SayHello(name string) {
	fmt.Println("SayHello: ", name)
}

func sayGoodbye(name string) {
	fmt.Println("sayGoodbye: ", name)
}
