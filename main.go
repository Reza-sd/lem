package main

import (
	"fmt"
	"main/util"
)

func main() {
	//var tunnelArr []string

	tunnelArr := []string{"0-4", "0-6", "1-3", "4-3"}
	util.InitGraph(tunnelArr)

	fmt.Println("------")

}
