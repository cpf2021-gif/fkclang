```text                              
                   ,--.              
    ,---,.     ,--/  /|    ,----..   
  ,'  .' |  ,---,': / '   /   /   \  
,---.'   |  :   : '/ /   |   :     : 
|   |   .'  |   '   ,    .   |  ;. / 
:   :  :    '   |  /     .   ; /--`  
:   |  |-,  |   ;  ;     ;   | ;     
|   :  ;/|  :   '   \    |   : |     
|   |   .'  |   |    '   .   | '___  
'   :  '    '   : |.  \  '   ; : .'| 
|   |  |    |   | '_\.'  '   | '/  : 
|   :  \    '   : |      |   :    /  
|   | ,'    ;   |,'       \   \ .'   
`----'      '---'          `---`                     
```

# Introduction
Fkclang is a simple interpreter and programming language implemented in Go.

# Example
recursion function
```
cong fib = fk(n) {
  if (n < 2) {
    -> n;
  } else {
    -> fib(n - 1) + fib(n - 2);
  }
}
```
high-level function
```
cong twice = fk(f, x) {
  -> f(f(x));
}

cong addTwo = fk(x) {
  -> x + 2;
}

twice(addTwo, 2); // -> 6
```

# Usage
You can run fkclang in two ways:
- Run the interpreter directly
```bash
go run main.go
```
- Run the interpreter with file
```bash
go run main.go xx.fkc
```
in the xx.fkc file, you can import other modules. For example:
```
import "std.fkc"
import "test.fkc"

// your code
```
It will run all the files in order, and the result will be printed to the terminal.You can continue to input code in the terminal after the file is executed.

# TODO
- [X] import module

# Reference
- [Writing An Interpreter In Go](https://interpreterbook.com/)