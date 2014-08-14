package main

import handler "github.com/tgandrews/mailbag/incominghandler"

func main() {
	handler := handler.IncomingHandler{}
	handler.Run()
}
