package runtime

import "fmt"
import "io"
import "os"
import "piled/opcode"

var writer io.Writer

type VM struct {
	code   []int
	ip     int
	stack  []int
	stdout io.Writer
}

func NewVM(code []int) *VM {
	return &VM{
		code:   code,
		ip:     0,
		stack:  make([]int, 0),
		stdout: os.Stdout,
	}
}

func (vm *VM) Run() {
	for vm.ip < len(vm.code) {
		op := vm.code[vm.ip]
		vm.ip++

		switch op {
		case opcode.PUSH:
			val := vm.code[vm.ip]
			vm.ip++
			vm.push(val)
		case opcode.ADD:
			b := vm.pop()
			a := vm.pop()
			vm.push(a + b)
		case opcode.SUB:
			b := vm.pop()
			a := vm.pop()
			vm.push(a - b)
		case opcode.GT:
			b := vm.pop()
			a := vm.pop()
			var v int;
			if a > b {
				v = 1
			} else {
				v = 0
			}
			vm.push(v)
		case opcode.LT:
			b := vm.pop()
			a := vm.pop()
			var v int;
			if a < b {
				v = 1
			} else {
				v = 0
			}
			vm.push(v)
		case opcode.EQ:
			b := vm.pop()
			a := vm.pop()
			var v int;
			if a == b {
				v = 1
			} else {
				v = 0
			}
			vm.push(v)
		case opcode.PRINT:
			a := vm.pop()
			vm.Println(a)
		}
	}
}

func (vm *VM) Println(a ...any) {
	fmt.Fprintln(vm.stdout, a...)
}

func (vm *VM) push(val int) {
	vm.stack = append(vm.stack, val)
}

func (vm *VM) pop() int {
	if len(vm.stack) == 0 {
		panic("RUNTIME ERROR: STACK UNDERFLOW")
	}
	index := len(vm.stack) - 1
	val := vm.stack[index]
	vm.stack = vm.stack[:index]
	return val
}
