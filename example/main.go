package main

import (
	"log/util"
	"time"
)

func main() {
	f1 := util.NewLogger(util.Info, "./", "xx1.log", 10*1024)
	//a := "test"
	f1.Info("a")
	time.Sleep(1 * time.Second)

}
