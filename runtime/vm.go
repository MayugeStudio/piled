package runtime

import "fmt"
import "io"
import "os"

// TODO: Delete it
var writer io.Writer

type VM struct {
	code   []OPCode
	ip     int
	stack  []int
	stdout io.Writer
}

func NewVM(code []OPCode) *VM {
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
		case PUSH:
			val := int(vm.code[vm.ip])
			vm.ip++
			vm.push(val)
		case ADD:
			b := vm.pop()
			a := vm.pop()
			vm.push(a + b)
		case SUB:
			b := vm.pop()
			a := vm.pop()
			vm.push(a - b)
		case GT:
			b := vm.pop()
			a := vm.pop()
			var v int
			if a > b {
				v = 1
			} else {
				v = 0
			}
			vm.push(v)
		case LT:
			b := vm.pop()
			a := vm.pop()
			var v int
			if a < b {
				v = 1
			} else {
				v = 0
			}
			vm.push(v)
		case EQ:
			b := vm.pop()
			a := vm.pop()
			var v int
			if a == b {
				v = 1
			} else {
				v = 0
			}
			vm.push(v)
		case PRINT:
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
