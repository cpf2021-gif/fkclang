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
	entryFile   = "main.fkc"
	congRegex   = `cong\s+([a-zA-Z_][a-zA-Z0-9_]*)\s*=`
)

// 记录已经导入的文件
var fileSet = make(map[string]struct{})

func Start(in io.Reader, out io.Writer, fileNames ...string) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	if len(fileNames) > 1 {
		printErrors("Command", out, []string{"only one file can be executed at a time"})
	} else if len(fileNames) == 1 && fileNames[0] != entryFile {
		printErrors("Command", out, []string{"the entry file must be main.fkc"})
	} else {
		ExecFiles(fileNames, out, env)
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
			fmt.Print("...")
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
		// 防止重复导入
		if _, ok := fileSet[fileName]; ok {
			continue
		}
		fileSet[fileName] = struct{}{}

		// 读取文件
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

		// std.fkc => std.
		if fileName == entryFile {
			onceExec(out, userInput, env)
		} else {
			prefix := strings.Split(fileName, ".")[0] + "."
			onceExec(out, userInput, env, prefix)
		}
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
func onceExec(out io.Writer, input string, env *object.Environment, prefix ...string) {
	// 词法分析
	// 将input中所有的变量名替换为prefix+变量名
	// cong a = 1 => cong prefix.a = 1
	srcInput := input
	if len(prefix) != 0 {
		congPattern := regexp.MustCompile(congRegex)
		// 找到所有的cong语句
		matches := congPattern.FindAllStringSubmatch(srcInput, -1)
		variableMap := make(map[string]string)

		for _, match := range matches {
			oldVar := match[1]
			newVar := prefix[0] + oldVar
			variableMap[oldVar] = newVar
		}

		for oldVar, newVar := range variableMap {
			srcInput = strings.ReplaceAll(srcInput, oldVar, newVar)
		}
	}

	l := lexer.New(srcInput)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
		return
	}

	var evaluated object.Object
	if len(prefix) != 0 {
		evaluated = evaluator.Eval(program, env)
	} else {
		evaluated = evaluator.Eval(program, env)
	}

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
