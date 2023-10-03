package main

import (
	"fkclang/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}

	// 读取命令行参数
	var fileNames []string
	if len(os.Args) > 1 {
		fileNames = os.Args[1:]
	}

	fmt.Printf("Hello %s! This is the FkCongLang programming language!\n",
		u.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout, fileNames...)
}
