package main

import (
	"fmt"
	"os"

	"./log"
)

func main() {
	stdObj := log.NewLogger(os.Stdout)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	stdObj.Info("hahaha")
	stdObj.Info("xixixi")

	fileObj, err := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Println("Error Occurred while open the file.")
		return
	}

	logObj := log.NewLogger(fileObj)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	logObj.Info("hahaha")
	logObj.Info("xixixix")
}
