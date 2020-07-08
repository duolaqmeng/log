package util

import (
	"fmt"
	"path"
	"runtime"
	"strconv"
)

func getCallerInfo()(funcName,file string ,line int){
	//pc uintptr, file string, line int, ok bool
	caller, file, line, ok := runtime.Caller(3)
	if !ok{
		return
	}
	funcName=runtime.FuncForPC(caller).Name()
	funcName=path.Base(funcName)
	file=path.Base(file)
	fmt.Println(funcName,file,line)
	return funcName ,file,line
}

func getLevelStr(i int) string{
	levelStr:=strconv.Itoa(i)
	return levelStr
}
