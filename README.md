# Piled

A simple stack based programming language implemented in Golang.
There is no memory concept in piled but only stack.

> [!WARNING]
> This is not serious project. Just a recreational one.

## Example

### Arithmetic Operators
```
4 2 + print
4 2 - print
4 2 * print
4 2 / print
```

Piled supports basic arithmetic operators such as +, -, *, /.

The output of the above program should be
```
6
2
8
2
```

### Conditional expressions

Piled supports conditionals.

```
100 0 > if {
    255 print
} else {
    0 print
}
```

Since 100 is greater than 0, the above Piled program should output 255. 

### Loop

Piled supports loops.

```
10 while dup 0 > {
    dup print
    1 -
}
```

The above program should output
```
10
9
8
7
..
1
```

The sequence of words between `while` and `{` are considered the condition and
they are executed at each step of loop.

## Usage

#### Run a program from a file

```console
./piled compile <filename>
./piled run     <filename>
```

#### REPL

```console
./piled
```

Just compile this project and run `./piled`.

