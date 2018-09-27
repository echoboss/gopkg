package main

import (
	"flag"
	"time"
	"fmt"
)

var name = flag.String("name","","-name=cola")

var age int

func main() {
	flag.IntVar(&age,"age",0,"-age=10")
	t := flag.Duration("spent",time.Second,"-spent=1s")

	flag.Parse()
	flag.PrintDefaults()

	fmt.Println(flag.Parsed())
	fmt.Println(*t)

	fmt.Println(flag.Args())
	fmt.Println(flag.Arg(0))
	fmt.Println(flag.NArg())
	fmt.Println(flag.NFlag())

}
