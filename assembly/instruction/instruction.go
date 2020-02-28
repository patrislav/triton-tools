package instruction

import (
	"strconv"

	"github.com/patrislav/triton-tools/bct"
)

type Instruction interface {
	TryteWidth() uint
	Encode() []bct.Trybble
}

type Simple uint8

func (i Simple) TryteWidth() uint { return 1 }
func (i Simple) Encode() []bct.Trybble {
	if i.isSupportOnly() {
		return []bct.Trybble{0, bct.Trybble(i)}
	} else {
		return []bct.Trybble{bct.Trybble(i), 0}
	}
}
func (i Simple) String() string {
	imap := map[Simple]string{
		Load: "load",
		Stor: "stor",
		Nop:  "nop",
		Ret:  "ret",
		Neg:  "neg",
		Swap: "swap",
		Add:  "add",
		Drop: "drop",
		Dup:  "dup",
		Rot:  "rot",
		Shr:  "shr",
		Shl:  "shl",
		Irq:  "irq",
		Jmp:  "jmp",
		Call: "call",
		Bmi:  "bmi",
		Bz:   "bz",
		Bpl:  "bpl",
	}
	v, ok := imap[i]
	if !ok {
		return string(i)
	}
	return v
}
func (i Simple) isDominantOnly() bool {
	for _, x := range dominants {
		if i == x {
			return true
		}
	}
	return false
}
func (i Simple) isSupportOnly() bool {
	for _, x := range supports {
		if i == x {
			return true
		}
	}
	return false
}

const (
	LitTryte = 0b10_10_00
	LitWord  = 0b10_10_01
)

const (
	Stor   Simple = 0b10_00_10
	Isu    Simple = 0b10_00_00
	Load   Simple = 0b10_00_01
	IncDec Simple = 0b10_01_10
	MaxMin Simple = 0b10_01_00
	IstIsf Simple = 0b10_01_01
	Neg    Simple = 0b00_00_10
	Nop    Simple = 0b00_00_00
	Add    Simple = 0b00_00_01
	Rot    Simple = 0b00_01_10
	Swap   Simple = 0b00_01_00
	Dup    Simple = 0b00_01_01
	Shl    Simple = 0b01_10_10
	Drop   Simple = 0b01_10_00
	Shr    Simple = 0b01_10_01

	Jmp  Simple = 0b01_00_00
	Irq  Simple = 0b10_10_10
	Call Simple = 0b01_00_01
	Bmi  Simple = 0b01_01_10
	Bz   Simple = 0b01_01_00
	Bpl  Simple = 0b01_01_01
	Ret  Simple = 0b01_00_10
)

var dominants = []Simple{Stor, Isu, Load, IncDec, MaxMin, IstIsf}
var supports = []Simple{Jmp, Call, Bmi, Bz, Bpl, Ret, Irq}

type PushAbsRef struct {
	Label string
	Addr  int
}

func (i *PushAbsRef) TryteWidth() uint { return 3 }
func (i *PushAbsRef) Encode() []bct.Trybble {
	w := bct.WordFromInt(int(i.Addr))
	return []bct.Trybble{
		0, LitWord, // PUSH WORD
		w.Hi.Hi, w.Hi.Lo, // hi tryte
		w.Lo.Hi, w.Lo.Lo, // lo tryte
	}
}

type PushRelRef struct {
	Label string
	Delta int
}

func (i *PushRelRef) TryteWidth() uint { return 2 }
func (i *PushRelRef) Encode() []bct.Trybble {
	t := bct.TryteFromInt(i.Delta)
	return []bct.Trybble{
		0, LitTryte, // PUSH TRYTE
		t.Hi, t.Lo, // Data
	}
}

type PushWord struct {
	Data bct.Word
}

func (i PushWord) TryteWidth() uint { return 3 }
func (i PushWord) Encode() []bct.Trybble {
	hi := i.Data.Hi.Trybbles()
	lo := i.Data.Lo.Trybbles()
	return []bct.Trybble{
		0, LitWord, // PUSH WORD
		hi[0], hi[1], // hi tryte
		lo[0], lo[1], // lo tryte
	}
}

func NewPushWordFromDecimal(s string) PushWord {
	v, err := strconv.Atoi(s)
	if err != nil {
		v = 0
	}
	return PushWord{Data: bct.WordFromInt(v)}
}

func NewPushWordFromNonary(val string) PushWord {
	// TODO: nonary strings
	return PushWord{Data: bct.WordFromInt(0)}
}

func NewPushWordFromTernary(val string) PushWord {
	// TODO: ternary strings
	return PushWord{Data: bct.WordFromInt(0)}
}

func NewPushWord(v int) PushWord {
	return PushWord{Data: bct.WordFromInt(v)}
}

type PushTryte struct {
	Data bct.Tryte
}

func (i PushTryte) TryteWidth() uint { return 2 }
func (i PushTryte) Encode() []bct.Trybble {
	trybbles := i.Data.Trybbles()
	return []bct.Trybble{
		0, LitTryte, // PUSH TRYTE
		trybbles[0], trybbles[1], // Data
	}
}

func NewPushTryte(v int) PushTryte {
	if v > 364 || v < -364 {
		v = 0
	}
	return PushTryte{Data: bct.TryteFromInt(v)}
}

func NewPushTryteFromDecimal(s string) PushTryte {
	v, err := strconv.Atoi(s)
	if err != nil {
		v = 0
	}
	return PushTryte{Data: bct.TryteFromInt(v)}
}

func NewPushTryteFromNonary(val string) PushTryte {
	// TODO: nonary strings
	return PushTryte{Data: bct.TryteFromInt(0)}
}

func NewPushTryteFromTernary(val string) PushTryte {
	// TODO: ternary strings
	return PushTryte{Data: bct.TryteFromInt(0)}
}

type PushTrybble struct {
	Data bct.Trybble
}

func (i PushTrybble) TryteWidth() uint { return 1 }
func (i PushTrybble) Encode() []bct.Trybble {
	return []bct.Trybble{
		LitTryte, i.Data, // PUSH TRYBBLE
	}
}

func NewPushTrybble(v int) PushTrybble {
	if v > 13 || v < -13 {
		v = 0
	}
	return PushTrybble{Data: bct.IntToTrybble(v)}
}
