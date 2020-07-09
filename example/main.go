package main

import (
	"log/util"
	"time"
)

func main() {
	f1 := util.NewLogger(util.Info, "./", "xx1.log")
	//a := "test"
	f1.Info("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	time.Sleep(1 * time.Second)

}
