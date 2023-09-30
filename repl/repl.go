package repl

import (
	"bufio"
	"fkclang/lexer"
	"fkclang/parser"
	"fmt"
	"io"
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
'---'      	'---'          '--'          
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Println(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}
