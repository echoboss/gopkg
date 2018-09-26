package main

import (
	"os/user"
	"fmt"
)

func main() {
	u, _ := user.Current()
	fmt.Printf("%+v\n",u)

	_, err :=user.Lookup("demo")
	fmt.Println(err)

	user, err := user.LookupId("501")
	fmt.Println(user,err)
}
