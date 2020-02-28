package assembly

import (
	"strconv"

	"github.com/patrislav/triton-tools/assembly/instruction"
	"github.com/patrislav/triton-tools/assembly/lexer"
)

type label struct {
	name   string
	index  int
	offset uint
}

type Parser struct {
	tokens []lexer.Token
	pos    int
}

func NewParser(tokens []lexer.Token) *Parser {
	p := &Parser{tokens: tokens}
	return p
}

func (p *Parser) ParseProgram(origin int) *Program {
	prog := &Program{origin: origin}

	var offset uint
	appendInstr := func(i instruction.Instruction) {
		prog.instructions = append(prog.instructions, i)
		offset += i.TryteWidth()
	}
	for curToken := p.nextToken(); curToken.Kind != lexer.TokenEOF; curToken = p.nextToken() {
		switch curToken.Kind {
		case lexer.TokenDirective:
			switch curToken.Val {
			case "word":

			}
		case lexer.TokenLabel:
			l := label{name: curToken.Val, index: len(prog.instructions), offset: offset}
			prog.labels = append(prog.labels, l)
		case lexer.TokenIdent:
			// TODO: error handling
			instr, _ := parseInstruction(curToken)
			appendInstr(instr)
		case lexer.TokenAddressAbsolute:
			instr := &instruction.PushAbsRef{Label: curToken.Val}
			appendInstr(instr)
		case lexer.TokenAddressRelative:
			instr := &instruction.PushRelRef{Label: curToken.Val}
			appendInstr(instr)
		case lexer.TokenDecimal:
			v, err := strconv.Atoi(curToken.Val)
			if err != nil {
				v = 0
			}
			var instr instruction.Instruction
			switch {
			case v > 364 || v < -364:
				instr = instruction.NewPushWord(v)
			case v > 13 || v < -13:
				instr = instruction.NewPushTryte(v)
			default:
				instr = instruction.NewPushTrybble(v)
			}
			appendInstr(instr)
		case lexer.TokenNonary:
			instr := instruction.NewPushTryteFromNonary(curToken.Val)
			appendInstr(instr)
		case lexer.TokenTernary:
			instr := instruction.NewPushTryteFromTernary(curToken.Val)
			appendInstr(instr)
		}
	}
	return prog
}

func (p *Parser) nextToken() lexer.Token {
	if p.pos >= len(p.tokens) {
		return lexer.Token{Kind: lexer.TokenEOF}
	}
	curToken := p.tokens[p.pos]
	p.pos++
	return curToken
}

var instrMap = map[string]instruction.Instruction{
	"load": instruction.Load,
	"stor": instruction.Stor,
	"nop":  instruction.Nop,
	"ret":  instruction.Ret,
	"neg":  instruction.Neg,
	"swap": instruction.Swap,
	"add":  instruction.Add,
	"drop": instruction.Drop,
	"dup":  instruction.Dup,
	"rot":  instruction.Rot,
	"irq":  instruction.Irq,
	"jmp":  instruction.Jmp,
	"call": instruction.Call,
	"bmi":  instruction.Bmi,
	"bz":   instruction.Bz,
	"bpl":  instruction.Bpl,
}

func parseInstruction(tok lexer.Token) (instruction.Instruction, error) {
	i, ok := instrMap[tok.Val]
	if !ok {
		// TODO: error handling
		return instruction.Nop, nil
	}
	return i, nil
}
