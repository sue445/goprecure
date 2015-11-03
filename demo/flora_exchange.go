package main

import (
	"fmt"
	"github.com/k0kubun/pp"
	"github.com/sue445/goprecure"
	"time"
)

func sleep()  {
	time.Sleep(1 * time.Second)
}

func main() {
	fmt.Println(">", "flora := goprecure.NewCureFlora()")
	flora := goprecure.NewCureFlora()
	fmt.Println()
	sleep()

	flora.PrintLine = false

	fmt.Println(">", "flora.Transform()")
	flora.Transform()
	fmt.Println()
	sleep()

	flora.PrintLine = true

	fmt.Println(">", "flora.Exchange(goprecure.Flora)")
	flora.Exchange(goprecure.Flora)
	fmt.Println()
	sleep()

	fmt.Println(">", "_, err := flora.Exchange(goprecure.Mermaid)")
	_, err := flora.Exchange(goprecure.Mermaid)
	pp.Println(err)
	fmt.Println()
}
