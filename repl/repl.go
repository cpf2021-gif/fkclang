package repl

import (
	"bufio"
	"fkclang/evaluator"
	"fkclang/lexer"
	"fkclang/object"
	"fkclang/parser"
	"fmt"
	"io"
	"strings"
)

const PROMPT = ">> "

const FKC = `
                   ,--.              
    ,---,.     ,--/  /|    ,----..   
  ,'  .' |  ,---,': / '   /   /   \  
,---.'   |  :   : '/ /   |   :     : 
|   |   .'  |   '   ,    .   |  ;. / 
:   :  :    '   |  /     .   ; /---
|   |  |-,  |   ;  ;     ;   | ;
|   :  ;/|  :   '   \    |   : |     
|   |   .'  |   |    '   .   | '___  
'   :  '    '   : |.  \  '   ; : .'| 
|   |  |    |   | '_\.'  '   | '/  :
|   :  \    '   : |      |   :    /  
|   | ,'    ;   |,'       \   \ .'
 ---         ---           ---
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

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
			_, _ = io.WriteString(out, FKC)
			_, _ = io.WriteString(out, "Whoops! We have encountered some errors here!\n")
			_, _ = io.WriteString(out, "evaluate errors:\n")
			_, _ = io.WriteString(out, "\t"+evaluated.Inspect())
			_, _ = io.WriteString(out, "\n")
		} else {
			_, _ = io.WriteString(out, evaluated.Inspect())
			_, _ = io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	_, _ = io.WriteString(out, FKC)
	_, _ = io.WriteString(out, "Whoops! We have encountered some errors here!\n")
	_, _ = io.WriteString(out, "parser errors:\n")
	for _, msg := range errors {
		_, _ = io.WriteString(out, "\t"+msg+"\n")
	}
}
