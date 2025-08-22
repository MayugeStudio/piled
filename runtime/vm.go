package runtime

import (
	"fmt"
	"io"
	"os"
)

// ReadBytecodeFile returns a slice of Instruction and error
// This function is helper function
func ReadBytecodeFile(path string) ([]Instruction, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	c := make([]Instruction, 0, 1024)
	for ip := 0; ip < len(raw); ip++ {
		kind := InstructionKind(raw[ip])
		switch kind {
		case PUSH, JMP, JMPIF:
			ip++
			if ip >= len(raw) {
				return nil, fmt.Errorf("error - reading bytecode: invalid bytecode was provided")
			}
			value := int(raw[ip])
			c = append(c, Instruction{Kind: kind, Args: []int{value}})
		default:
			c = append(c, Instruction{Kind: kind, Args: []int{}})
		}
	}

	return c, nil
}

// VM contains piled Virtual Machine state
type VM struct {
	code   []Instruction
	ip     int
	stack  []int
	stdout io.Writer
}

// NewVM is constructor for piled Virtual Machine
func NewVM(code []Instruction) *VM {
	return &VM{
		code:   code,
		ip:     0,
		stack:  make([]int, 0),
		stdout: os.Stdout,
	}
}

// Run start emulating a slice of Instruction on VM
func (vm *VM) Run() {
	for vm.ip < len(vm.code) {
		inst := vm.code[vm.ip]
		vm.ip++
		switch inst.Kind {
		case PUSH:
			val := inst.Args[0]
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
			addr := inst.Args[0]
			vm.ip = addr
		case JMPIF:
			// TODO: JMPIF need to be jump if cond is true but now it jumps if cond is false

			// address is at next opcode
			addr := inst.Args[0]
			cond := vm.pop()
			if cond == 0 { // false
				vm.ip = addr
			} else { // true
				// fallthrough
			}
		case PRINT:
			a := vm.pop()
			vm.Println(a)
		case DUP:
			a := vm.pop()
			vm.push(a)
			vm.push(a)

		case SWAP:
			b := vm.pop()
			a := vm.pop()
			vm.push(b)
			vm.push(a)
		case ROT:
			c := vm.pop()
			b := vm.pop()
			a := vm.pop()

			vm.push(b)
			vm.push(c)
			vm.push(a)
		case DROP:
			vm.pop()

		case NOP: // do nothing
		default:
			panic("RUNTIME ERROR: INVALID INSTRUCTION PROVIDED")
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
