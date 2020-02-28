package assembly

import (
	"reflect"
	"testing"

	"github.com/patrislav/triton-tools/assembly/instruction"
	"github.com/patrislav/triton-tools/assembly/lexer"
	"github.com/patrislav/triton-tools/bct"
)

func TestParser_ParseProgram(t *testing.T) {
	input := []lexer.Token{
		{Kind: lexer.TokenIdent, Val: "nop"},
		{Kind: lexer.TokenLabel, Val: "main"},
		{Kind: lexer.TokenIdent, Val: "nop"},
		{Kind: lexer.TokenDecimal, Val: "123"},
		{Kind: lexer.TokenDecimal, Val: "-123"},
		{Kind: lexer.TokenComment, Val: "comment"}, // comments are ignored
		{Kind: lexer.TokenLabel, Val: "hello"},
		{Kind: lexer.TokenIdent, Val: "jmp"},
		{Kind: lexer.TokenAddressAbsolute, Val: "main"},
		{Kind: lexer.TokenAddressRelative, Val: "hello"},
		// TODO: ternary and nonary literals
		// {Kind: lexer.TokenTernary, Val: "T1001T"},
		// {Kind: lexer.TokenNonary, Val: "1ADB3C"},
	}
	p := NewParser(input)
	prog := p.ParseProgram(0)

	t.Run("instructions", func(t *testing.T) {
		want := []instruction.Instruction{
			instruction.Nop,
			instruction.Nop,
			instruction.PushTryte{Data: bct.TryteFromInt(123)},
			instruction.PushTryte{Data: bct.TryteFromInt(-123)},
			instruction.Jmp,
			&instruction.PushAbsRef{Label: "main"},
			&instruction.PushRelRef{Label: "hello"},
			// TODO: ternary and nonary literals
			// pushWordInstruction{data: bct.WordFromInt(-160)},
			// pushWordInstruction{data: bct.WordFromInt(49434)},
		}
		for i, instr := range prog.instructions {
			if len(want) <= i {
				t.Fatalf("The resulting Program has too many instructions")
			}
			if !reflect.DeepEqual(instr, want[i]) {
				t.Errorf("got = %v, want = %v", instr, want[i])
			}
		}
		if len(prog.instructions) < len(want) {
			t.Fatalf("The resulting Program does not have enough instructions")
		}
	})

	t.Run("labels", func(t *testing.T) {
		want := []label{
			{name: "main", index: 1, offset: 1},
			{name: "hello", index: 4, offset: 6},
		}
		for i, lab := range prog.labels {
			if len(want) <= i {
				t.Fatalf("The resulting Program has too many labels")
			}
			if !reflect.DeepEqual(lab, want[i]) {
				t.Errorf("got = %v, want = %v", lab, want[i])
			}
		}
		if len(prog.labels) < len(want) {
			t.Fatalf("The resulting Program does not have enough labels")
		}
	})

}
