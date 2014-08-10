package main

import "fmt"
import handler "github.com/tgandrews/mailbag/incominghandler"

func main() {
	fmt.Printf("Hello world")
	handler := handler.IncomingHandler{}
	handler.Run()
}
