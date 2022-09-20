package main

import (
	core "CommonCheckIn/core"
	"flag"
	"fmt"
)

var testFlg = flag.Bool("test", false, "false")

func init() {
	flag.Parse()
}

func main() {
	fmt.Println("Start!")
	core.Run(*testFlg)
	fmt.Println("End!")
}
