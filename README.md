# Monkey programming language

## Introduction

Created a new programming language called Monkey.
Monkey was implemented twice using two different approaches: 
- 1st - interpreter approach: Pratt parsing -> Tree-walking interpreting
- 2nd - compiler + VM: Pratt parsing -> Compiling to bytecode -> VM

Did a benchmark which showed that the compiler approach offers **three times faster** execution time than that of the other approach.  

```bash
hxt365@Truongs-Air monkey % go run benchmark/main.go -engine="vm"
>> engine=vm, result=9227465, duration=4.354578s
hxt365@Truongs-Air monkey % go run benchmark/main.go -engine="eval"
>> engine=eval, result=9227465, duration=15.337929584s
```

## Resources

- Writing An Interpreter In Go book by Thorsten Ball
- Writing A Compiler In Go book by Thorsten Ball