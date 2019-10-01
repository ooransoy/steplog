package main

import (
	sl "github.com/ooransoy/steplog"
)

var log = sl.NewSteplogger()

func init() {
	sl.StripedPrint = true
}

func main() {
	log.Open("Hi!")
	log.Print("Hello!")
	log.Close()
	log.Open("Hola!")
	log.Print("Aloha!")
	log.Print("Goodbye.")
}
