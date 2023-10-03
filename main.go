package main

import (
	"fkclang/repl"
	"fmt"
	"os"
	"os/user"
	"strings"
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

	// 检查文件后缀名
	checkExtension(fileNames)

	fmt.Printf("Hello %s! This is the FkCongLang programming language!\n",
		u.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout, fileNames...)
}

func checkExtension(fileNames []string) {
	for _, fileName := range fileNames {
		if !strings.HasSuffix(fileName, ".fkc") {
			panic("file extension must be .fkc")
		}
	}
}
