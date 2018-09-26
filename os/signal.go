package main

import (
	"os"
	"os/signal"
	"fmt"
)

func main() {
	sign := make(chan os.Signal)
	signal.Notify(sign,os.Interrupt)
	s := <- sign
	fmt.Println("I catch signal: ", s.String())
}