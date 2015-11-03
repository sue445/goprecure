package main

import (
	"fmt"
	"github.com/k0kubun/pp"
	"github.com/sue445/goprecure"
	"time"
)

func sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	fmt.Println(">", "flora := goprecure.NewCureFlora()")
	flora := goprecure.NewCureFlora()
	fmt.Println()
	sleep()

	fmt.Println(">", "flora.Name()")
	pp.Println(flora.Name())
	fmt.Println()
	sleep()

	fmt.Println(">", "flora.Transform()")
	flora.Transform()
	fmt.Println()
	sleep()

	fmt.Println(">", "flora.Name()")
	pp.Println(flora.Name())
}
