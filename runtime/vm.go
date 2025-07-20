package runtime

import "fmt"
import "io"
import "os"
import "piled/opcode"

var writer io.Writer

func init() {
	writer = os.Stdout
}

func vmPrint(a ...any) {
	fmt.Fprintln(writer, a...)
}

type VM struct {
	code  []int
	ip    int
	stack []int
}

func NewVM(code []int) *VM {
	return &VM{code: code, ip: 0, stack: make([]int, 0)}
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
		case opcode.PRINT:
			vmPrint(vm.pop())
		}
	}
}

func (vm *VM) push(val int) {
	vm.stack = append(vm.stack, val)
}

func (vm *VM) pop() int {
	if len(vm.stack) == 0 {
		panic("RUNTIME ERROR")
	}
	index := len(vm.stack) - 1
	val := vm.stack[index]
	vm.stack = vm.stack[:index]
	return val
}
