package codegen

import (
	"fmt"
	"os"
	"piled/compiler/ir"
	"piled/runtime"
)

// GenerateProgram generate instruction from source code.
// TODO: Make GenerateProgram output program into file directly
//
//	ex. func GenerateProgram(output strings.Builder, [])
func GenerateProgram(body []ir.Op) []runtime.Instruction {
	output := make([]runtime.Instruction, 0)
	labelMap := make(map[int]int)
	// Prefinding Labels
	for ip := 0; ip < len(body); ip++ {
		switch v := body[ip].(type) {
		case *ir.Label:
			{
				labelMap[v.Label] = ip
			}
		}
	}

	for ip := 0; ip < len(body); ip++ {
		op := body[ip]
		switch v := op.(type) {
		case *ir.Number:
			emit(&output, runtime.Instruction{Kind: runtime.PUSH, Args: []int{v.Value}})
		case *ir.Bind:
			fmt.Println("Binding is not supported yet")
		case *ir.Binop:
			switch v.Bkind {
			case ir.Add:
				emit(&output, runtime.Instruction{Kind: runtime.ADD, Args: nil})
			case ir.Sub:
				emit(&output, runtime.Instruction{Kind: runtime.SUB, Args: nil})
			case ir.Mul:
				emit(&output, runtime.Instruction{Kind: runtime.MUL, Args: nil})
			case ir.Div:
				emit(&output, runtime.Instruction{Kind: runtime.DIV, Args: nil})
			case ir.Mod:
				emit(&output, runtime.Instruction{Kind: runtime.MOD, Args: nil})
			case ir.And:
				emit(&output, runtime.Instruction{Kind: runtime.AND, Args: nil})
			case ir.Gt:
				emit(&output, runtime.Instruction{Kind: runtime.GT, Args: nil})
			case ir.Lt:
				emit(&output, runtime.Instruction{Kind: runtime.LT, Args: nil})
			case ir.Eq:
				emit(&output, runtime.Instruction{Kind: runtime.EQ, Args: nil})
			case ir.Or:
				emit(&output, runtime.Instruction{Kind: runtime.OR, Args: nil})
			case ir.Shl:
				emit(&output, runtime.Instruction{Kind: runtime.SHL, Args: nil})
			case ir.Shr:
				emit(&output, runtime.Instruction{Kind: runtime.SHR, Args: nil})
			}
		case *ir.Label:
			emit(&output, runtime.Instruction{Kind: runtime.NOP, Args: nil})
		case *ir.JmpLabel:
			labelAddr := labelMap[v.Label]
			emit(&output, runtime.Instruction{Kind: runtime.JMP, Args: []int{labelAddr}})
		case *ir.JmpIfNotLabel:
			labelAddr := labelMap[v.Label]
			emit(&output, runtime.Instruction{Kind: runtime.JMPIF, Args: []int{labelAddr}})
		case *ir.Print:
			emit(&output, runtime.Instruction{Kind: runtime.PRINT, Args: nil})
		case *ir.Dup:
			emit(&output, runtime.Instruction{Kind: runtime.DUP, Args: nil})
		case *ir.Swap:
			emit(&output, runtime.Instruction{Kind: runtime.SWAP, Args: nil})
		case *ir.Rot:
			emit(&output, runtime.Instruction{Kind: runtime.ROT, Args: nil})
		case *ir.Drop:
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
