package main

import "log/util"

func main() {
	f1 := util.NewLogger(util.Info, "./", "xx")
	a := "test"
	f1.Info("this is %s hello1111111111", a)

}
