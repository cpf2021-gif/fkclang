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
## Install
First, you need to install Go and set the environment variables. Then you can clone the project and build it.
```bash
go build
```
when you run the command, you will get an executable file named fkclang.
## Run
You can run fkclang in two ways:
- Run the interpreter directly
```bash
./fkclang
```
- Run the interpreter with file called main.fkc
```bash
./fkclang main.fkc
```
in the main.fkc file, you can import other modules. For example:
```
import "std.fkc"

// use std.fkc's sort function
print(std.sort([3, 2, 1], 1));
```
You can continue to input code in the terminal after the file is executed.

# TODO
- [X] import module

# Reference
- [Writing An Interpreter In Go](https://interpreterbook.com/)