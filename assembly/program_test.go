package assembly

import (
	"reflect"
	"testing"

	"github.com/patrislav/triton-tools/assembly/instruction"
	"github.com/patrislav/triton-tools/bct"
)

func TestProgram_ResolveAddresses(t *testing.T) {
	labels := []label{
		{name: "main", index: 1, offset: 1},
		{name: "hello", index: 4, offset: 6},
	}
	input := []instruction.Instruction{
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
	want := []instruction.Instruction{
		instruction.Nop,
		instruction.Nop,
		instruction.PushTryte{Data: bct.TryteFromInt(123)},
		instruction.PushTryte{Data: bct.TryteFromInt(-123)},
		instruction.Jmp,
		&instruction.PushAbsRef{Label: "main", Addr: 1},
		&instruction.PushRelRef{Label: "hello", Delta: -7},
	}
	prog := &Program{
		instructions: input,
		labels: labels,
	}
	if err := prog.ResolveAddresses(); err != nil {
		t.Fatalf("Unexpected error: %w", err)
	}
	for i := range prog.instructions {
		g, w := prog.instructions[i], want[i]
		if !reflect.DeepEqual(g, w) {
			t.Errorf("instruction %d: got = %v, want = %v", i, g, w)
		}
	}
}
