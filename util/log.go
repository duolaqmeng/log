package util

import (
	"fmt"
	"os"
	"time"
)

const (
	Debug  = iota
	Trace
	Info
	Warn
	Error
	Cirtal
)

type Logger struct {
	Level int
	logfilePath string
	logfileName string
	file *os.File
}

func NewLogger(level int,logfilePath, logfileName string) *Logger{
	f1:=&Logger{
		Level:       level,
		logfileName: logfileName,
		logfilePath: logfileName,
	}
	f1.initLogger()
	return f1
}


func(l *Logger)initLogger(){
	filepath:=fmt.Sprintf("%s%s",l.logfilePath,l.logfileName)
	file ,err:=os.OpenFile(filepath,os.O_CREATE|os.O_WRONLY|os.O_APPEND,0655)
	if err!=nil{
		panic("open file "+filepath+" error")
	}
	l.file=file
}

func (l *Logger)Debug(format string, a ...interface{}){
	if  l.Level>Debug{
		return
	}
	l.Common(format ,a...)
}
func (l *Logger)Trace(format string, a ...interface{}){
	if  l.Level>Trace{
		return
	}
	l.Common(format,a...)
}
func (l *Logger)Warn(format string, a ...interface{}){
	if  l.Level>Warn{
		return
	}
	l.Common(format,a...)
}

func (l *Logger)Error(format string, a ...interface{}){
	if  l.Level>Error{
		return
	}
	l.Common(format,a...)
}

func (l *Logger)Common(format string ,a... interface{}){
	timeStr:=time.Now().Format("[2006-01-02 15:04:05.000]")
	funcname, file, line := getCallerInfo()
	format=fmt.Sprintf("[%s][-%s-] [%s %s:%d] %s",
		timeStr,getLevelStr(l.Level),file,funcname,line,format)
	fmt.Fprintf(l.file,format,a...)
	fmt.Fprintln(l.file)
}




