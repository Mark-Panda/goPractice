package main

import (
	"fmt"
	"mylog"
)

func main()  {
	fmt.Println("sssssss")
	logger := mylog.NewFileLogger("Debug","./","xxx.log")
	sb := "官大吗"
	fmt.Println("llllllll")
	logger.Debug("%s是个跟", sb)
	logger.Info("这是一条测试Info")
	logger.Error("Error 这还少")
}