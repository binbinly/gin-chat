package main

import (
	"github.com/rs/xid"
)

func main() {
	println(xid.New().String())
	println(xid.New().String())
	println(xid.New().String())
	println(xid.New().String())
	println(xid.New().String())
	println(xid.New().String())
}
