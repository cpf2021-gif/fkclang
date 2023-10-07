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

# Features
- c-like syntax
- support integer, boolean, string, array, map, function type
- arithmetical expression
- if-else statement
- built-in functions
- high-level function
- import module
- closure

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

# Example
### declare variable and assign value
```text
// declare variable with initialization
cong a = 1;
// assign value
set a = 2;
set a = "hello world";
set a = [1, 2, 3];
set a = {a: 1, b: 2};
set a = fk(x) {
  -> x + 1;
};
```

### use index to get element
```text
cong a = [1, 2, 3];
print(a[0]); // -> 1
```

### built-in functions
```text
// print
print(1); // -> 1
// len
print(len([1, 2, 3])); // -> 3
```
to see more built-in functions, you can check the `evaluator/builtins.go` file.

### recursion function
```text
cong fib = fk(n) {
  if (n < 2) {
    -> n;
  } else {
    -> fib(n - 1) + fib(n - 2);
  }
}

fib(4); // -> 3
```
### high-level function
```text
cong twice = fk(f, x) {
  -> f(f(x));
}

cong addTwo = fk(x) {
  -> x + 2;
}

twice(addTwo, 2); // -> 6
```
to see more examples, you can check the `std.fkc` file.

# TODO
- [X] import module
- [ ] float number
- [ ] match syntax
- [ ] for loop (maybe difficult)

# Reference
- [Writing An Interpreter In Go](https://interpreterbook.com/)