package main

import (
	"fmt"
	"errors"
)

func main() {
	fmt.Printf("%T\n",errors.New("fuck error"))
	fmt.Printf("%T\n",fmt.Errorf("fuck error"))
}
