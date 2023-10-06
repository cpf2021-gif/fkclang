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
	"regexp"
	"strings"
)

const (
	PROMPT      = ">> "
	importRegex = `import\s+"([^"]+)"`
)

func Start(in io.Reader, out io.Writer, fileNames ...string) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	ExecFiles(fileNames, out, env)

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

		importFileNames, srcInput, errors := processInput(inputs)
		if len(errors) != 0 {
			printImportErrors(out, errors)
			continue
		}
		ExecFiles(importFileNames, out, env)
		onceExec(out, srcInput, env)
	}
}

func ExecFiles(fileNames []string, out io.Writer, env *object.Environment) {
	for _, fileName := range fileNames {
		// 读取文件并执行
		inputs, errors := readFile(fileName)
		if len(errors) != 0 {
			printImportErrors(out, errors)
			break
		}
		importFileNames, userInput, errors := processInput(inputs)
		if len(errors) != 0 {
			printImportErrors(out, errors)
			break
		}
		ExecFiles(importFileNames, out, env)
		onceExec(out, userInput, env)
	}
}

func readFile(fileName string) (src []string, errors []string) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, []string{"this module is not found"}
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	fileScanner := bufio.NewScanner(file)
	var inputs []string
	for fileScanner.Scan() {
		inputs = append(inputs, fileScanner.Text())
	}
	return inputs, nil
}

// 处理输入
// 将import和其他语句分开
func processInput(input []string) (importFilename []string, remainString string, ParseErrors []string) {
	var imports []string
	remain := ""
	importEnd := false
	for _, line := range input {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "import") {
			if importEnd {
				return nil, "", []string{"import must be at the top of the file"}
			}
			// import "std.fkc" -> std.fkc
			re := regexp.MustCompile(importRegex)
			matches := re.FindStringSubmatch(trimmed)

			if len(matches) > 1 {
				fileName := matches[1]
				imports = append(imports, fileName)
			} else {
				return nil, "", []string{`Error syntax of import, should be import "std.fkc"`}
			}

		} else {
			if trimmed != "" {
				if !importEnd {
					importEnd = true
				}
				remain += line + "\n"
			}
		}
	}

	return imports, remain, nil
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
