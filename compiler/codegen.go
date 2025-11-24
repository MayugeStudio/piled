package compiler

import (
	"fmt"
	"os"
	"piled/runtime"
)

// GenerateProgram generate instruction from source code.
// TODO: Make GenerateProgram output program into file directly
//
//	ex. func GenerateProgram(output strings.Builder, [])
func GenerateProgram(body []Op) []runtime.Instruction {
	output := make([]runtime.Instruction, 0)
	labelMap := make(map[int]int)
	// Prefinding Labels
	for ip := 0; ip < len(body); ip++ {
		switch v := body[ip].(type) {
		case *Label:
			{
				labelMap[v.Label] = ip
			}
		}
	}

	for ip := 0; ip < len(body); ip++ {
		op := body[ip]
		switch v := op.(type) {
		case *Number:
			emit(&output, runtime.Instruction{Kind: runtime.PUSH, Args: []int{v.Value}})
		case *Bind:
			fmt.Println("Binding is not supported yet")
		case *Binop:
			switch v.Bkind {
			case Add:
				emit(&output, runtime.Instruction{Kind: runtime.ADD, Args: nil})
			case Sub:
				emit(&output, runtime.Instruction{Kind: runtime.SUB, Args: nil})
			case Mul:
				emit(&output, runtime.Instruction{Kind: runtime.MUL, Args: nil})
			case Div:
				emit(&output, runtime.Instruction{Kind: runtime.DIV, Args: nil})
			case Mod:
				emit(&output, runtime.Instruction{Kind: runtime.MOD, Args: nil})
			case And:
				emit(&output, runtime.Instruction{Kind: runtime.AND, Args: nil})
			case Gt:
				emit(&output, runtime.Instruction{Kind: runtime.GT, Args: nil})
			case Lt:
				emit(&output, runtime.Instruction{Kind: runtime.LT, Args: nil})
			case Eq:
				emit(&output, runtime.Instruction{Kind: runtime.EQ, Args: nil})
			case Or:
				emit(&output, runtime.Instruction{Kind: runtime.OR, Args: nil})
			case Shl:
				emit(&output, runtime.Instruction{Kind: runtime.SHL, Args: nil})
			case Shr:
				emit(&output, runtime.Instruction{Kind: runtime.SHR, Args: nil})
			}
		case *Label:
			emit(&output, runtime.Instruction{Kind: runtime.NOP, Args: nil})
		case *JmpLabel:
			labelAddr := labelMap[v.Label]
			emit(&output, runtime.Instruction{Kind: runtime.JMP, Args: []int{labelAddr}})
		case *JmpIfNotLabel:
			labelAddr := labelMap[v.Label]
			emit(&output, runtime.Instruction{Kind: runtime.JMPIF, Args: []int{labelAddr}})
		case *Print:
			emit(&output, runtime.Instruction{Kind: runtime.PRINT, Args: nil})
		case *Dup:
			emit(&output, runtime.Instruction{Kind: runtime.DUP, Args: nil})
		case *Swap:
			emit(&output, runtime.Instruction{Kind: runtime.SWAP, Args: nil})
		case *Rot:
			emit(&output, runtime.Instruction{Kind: runtime.ROT, Args: nil})
		case *Drop:
			emit(&output, runtime.Instruction{Kind: runtime.DROP, Args: nil})
		default:
			fmt.Printf("CODE-GEN: unhandled op: %v\n", op)
		}
	}
	return output
}

func emit(output *[]runtime.Instruction, i runtime.Instruction) {
	*output = append(*output, i)
}

// Write output an array of instruction to specified filepath
func Write(path string, insts []runtime.Instruction) error {
	out := make([]byte, 0, 1024)
	for _, inst := range insts {
		out = append(out, byte(inst.Kind))
		if len(inst.Args) > 0 {
			for _, arg := range inst.Args {
				out = append(out, byte(arg))
			}
		}
	}
	return os.WriteFile(path, out, 0644)
}
