package main

import (
	"fmt"
	"time"
)

type user struct {
	name string
	age  int8
}

var u = user{
	name: "Ankur",
	age:  25,
}
var g = &u

func modifyUser(pu *user) {
	fmt.Println("modifyUser Received Value", pu)
	pu.name = "Anand"
}

func printUser(u <-chan *user) {
	time.Sleep(2 * time.Second)
	fmt.Println("printUser goRoutine called", <-u)
}

// channel 发送和接收元素的本质是 值的拷贝(通信共享内存)
func main() {
	c := make(chan *user, 5)
	c <- g
	fmt.Printf("value:%+v type:%T\n", g, g)
	// modify g
	g = &user{
		name: "Ankur Anand",
		age:  100,
	}
	go printUser(c)
	go modifyUser(g)
	time.Sleep(5 * time.Second)
	fmt.Println(g)
}
