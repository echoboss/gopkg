package main

import (
	"net/rpc"
	"net"
	"log"
	"net/rpc/jsonrpc"
)


type Rect struct{}

func (r *Rect) Area(p Params, resp *int) error {
	*resp = p.Height * p.Height
	return nil
}

func (r *Rect) Perimeter(p Params, resp *int) error {
	*resp = (p.Height + p.Width) * 2
	return nil
}

const port  = ":50001"

func main() {
	rect := new(Rect)
	rpc.Register(rect)
	l, err := net.Listen("tcp",port)
	if err != nil {
		panic(err)
	}
	log.Printf("server is running %s",port)
	for  {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("accept failed: %v",err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}

