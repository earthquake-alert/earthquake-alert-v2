package main

import "github.com/earthquake-alert/erarthquake-alert-v2/src"

var mode = "local"

func init() {
	src.Init(mode)
}

func main() {
	src.Server()
}
