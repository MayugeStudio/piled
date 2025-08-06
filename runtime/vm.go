package runtime

import "fmt"
import "io"
import "os"

// ReadBytecodeFile returns a slice of OPCode and error
// This function is helper function
func ReadBytecodeFile(path string) ([]OPCode, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	c := make([]OPCode, 0, 1024)
	for _, r := range raw {
		c = append(c, OPCode(r))
	}

	return c, nil
}

// VM contains piled Virtual Machine state
type VM struct {
	code   []OPCode
	ip     int
	stack  []int
	stdout io.Writer
}

// NewVM is constructor for piled Virtual Machine
func NewVM(code []OPCode) *VM {
	return &VM{
		code:   code,
		ip:     0,
		stack:  make([]int, 0),
		stdout: os.Stdout,
	}
}

// Run start emulating a slice of OPCode on VM
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
		case MUL:
			b := vm.pop()
			a := vm.pop()
			vm.push(a * b)
		case DIV:
			b := vm.pop()
			a := vm.pop()
			vm.push(a / b)
		case MOD:
			b := vm.pop()
			a := vm.pop()
			vm.push(a % b)
		case AND:
			b := vm.pop()
			a := vm.pop()
			vm.push(a & b)
		case OR:
			b := vm.pop()
			a := vm.pop()
			vm.push(a | b)
		case SHL:
			b := vm.pop()
			a := vm.pop()
			vm.push(a << b)
		case SHR:
			b := vm.pop()
			a := vm.pop()
			vm.push(a >> b)
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
		case JMP:
			addr := vm.code[vm.ip]
			vm.ip = int(addr)
		case JMPIF:
			// TODO: JMPIF need to be jump if cond is true but now it jumps if cond is false
			// address is at next opcode
			addr := vm.code[vm.ip]
			vm.ip++ // ensure vm ip is point at the next opcode.
			cond := vm.pop()
			if cond == 0 { // false
				vm.ip = int(addr)
			} else { // true
				// fallthrough
			}

		case PRINT:
			a := vm.pop()
			vm.Println(a)
		case NOP: // do nothing
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
