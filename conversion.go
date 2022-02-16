package main

import "fmt"

func main7() {
	var nilai32 int32 = 100000
	var nilai16 int16 = int16(nilai32)
	var nilai64 int64 = int64(nilai32)

	fmt.Println("nilai32:", nilai32)
	fmt.Println("nilai16:", nilai16)
	fmt.Println("nilai64:", nilai64)

	var name string = "Randi"
	var e = name[0]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println("e:", e)
	fmt.Println("eString:", eString)

}
