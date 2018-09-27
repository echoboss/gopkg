package main

import (
	"net/rpc/jsonrpc"
	"fmt"
)



func main() {
	rpc, err := jsonrpc.Dial("tcp",":50001")
	if err != nil {
		panic(err)
	}
	var resp int
	err = rpc.Call("Rect.Area",Params{Height:2,Width:3},&resp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Resp: %d\n",resp)

	err = rpc.Call("Rect.Perimeter",Params{Height:2,Width:3},&resp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Resp: %d\n",resp)


}

// {"id":1,"jsonrpc":"2.0","method":"Rect.Area","params":[{"height":2,"width":3}]}
