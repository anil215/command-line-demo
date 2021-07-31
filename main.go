package main

import (
	"fmt"

	"example.com/command/commandengine"
)

// go run example.com/command
func main() {
	executor := commandengine.NewCommandEngine()
	err := executor.Run()
	if err != nil {
		fmt.Println(fmt.Sprintf("err occured in running command line application, %s", err.Error()))
	}
}
