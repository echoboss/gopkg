package main

import (
	"os"
	"log"
)

func main() {
	f, _ := os.OpenFile("acc.log",os.O_CREATE|os.O_APPEND|os.O_WRONLY,0644)
	l := log.New(f,"acc_",log.Ltime|log.Lmicroseconds)

	l.Printf("send to money for %s","abc")
	l.SetPrefix("tx_")
	l.SetFlags(log.Llongfile)

}
