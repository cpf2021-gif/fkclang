package repl

import (
	"bufio"
	"fkclang/evaluator"
	"fkclang/lexer"
	"fkclang/object"
	"fkclang/parser"
	"fmt"
	"io"
	"os"
	"strings"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer, fileNames ...string) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for _, fileName := range fileNames {
		// 读取文件并执行
		func() {
			file, err := os.Open(fileName)
			if err != nil {
				panic(err)
			}
			defer func(file *os.File) {
				err := file.Close()
				if err != nil {
					panic(err)
				}
			}(file)
			fileScanner := bufio.NewScanner(file)
			var inputs []string
			for fileScanner.Scan() {
				inputs = append(inputs, fileScanner.Text())
			}
			userInput := strings.Join(inputs, "\n")
			onceExec(out, userInput, env)
		}()
	}

	for {
		var inputs []string
		fmt.Print(PROMPT)

		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				break
			}
			inputs = append(inputs, line)
			fmt.Print("..")
		}

		userInput := strings.Join(inputs, "\n")
		onceExec(out, userInput, env)
	}
}

// 执行一次
func onceExec(out io.Writer, input string, env *object.Environment) {
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
		return
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		if evaluated.Type() == object.ErrorObj {
			errors := []string{
				evaluated.Inspect(),
			}
			printRunTimeErrors(out, errors)
		} else {
			_, _ = io.WriteString(out, evaluated.Inspect()+"\n")
		}
	}
}

func printImportErrors(out io.Writer, errors []string) {
	printErrors("Import", out, errors)
}

func printRunTimeErrors(out io.Writer, errors []string) {
	printErrors("Runtime", out, errors)
}

func printParserErrors(out io.Writer, errors []string) {
	printErrors("Parser", out, errors)
}

func printErrors(errorType string, out io.Writer, errors []string) {
	_, _ = io.WriteString(out, "Whoops! We have encountered some errors here!\n")
	_, _ = io.WriteString(out, errorType+" errors:\n")
	for _, msg := range errors {
		_, _ = io.WriteString(out, "\t"+msg+"\n")
	}
}
