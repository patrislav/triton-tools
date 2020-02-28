package assembly

import (
	"fmt"

	"github.com/patrislav/triton-tools/assembly/instruction"
)

type Program struct {
	origin       int
	labels       []label
	instructions []instruction.Instruction
}

func (p *Program) ResolveAddresses() error {
	var pos uint
	for _, instr := range p.instructions {
		switch instr.(type) {
		case *instruction.PushAbsRef:
			i := instr.(*instruction.PushAbsRef)
			l := p.findLabel(i.Label)
			if l == nil {
				return fmt.Errorf("could not find label with name %q", i.Label)
			}
			i.Addr = p.origin + int(l.offset)
		case *instruction.PushRelRef:
			i := instr.(*instruction.PushRelRef)
			l := p.findLabel(i.Label)
			if l == nil {
				return fmt.Errorf("could not find label with name %q", i.Label)
			}
			i.Delta = int(l.offset) - int(pos + 1 + instr.TryteWidth())
		}
		pos += instr.TryteWidth()
	}
	return nil
}

func (p *Program) GetBytes() []byte {
	buf := make([]byte, 0)
	for _, instr := range p.instructions {
		for _, t := range instr.Encode() {
			buf = append(buf, byte(t))
		}
	}
	return buf
}

func (p *Program) findLabel(name string) *label {
	for _, l := range p.labels {
		if l.name == name {
			return &l
		}
	}
	return nil
}
