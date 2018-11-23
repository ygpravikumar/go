package main

import "fmt"
import "pluralsight/go/packages/model"

func main() {
	jumperList := model.GetJumperList()
	for _, jumper := range jumperList {
		fmt.Println(jumper.Jump())
	}
}
