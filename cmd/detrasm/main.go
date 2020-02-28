package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/patrislav/triton-tools/assembly/instruction"
	"github.com/patrislav/triton-tools/bct"
)

const origin = 9842

func main() {
	content, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Could not read file: %s", err)
	}

	for pos := 0; pos < len(content); {
		instr := bct.Tryte{Hi: bct.Trybble(content[pos]), Lo: bct.Trybble(content[pos+1])}
		o := 2

		if instr.Hi == instruction.LitTryte {
			t := bct.Tryte{Hi: bct.Trybble(0), Lo: instr.Lo}
			fmt.Printf("%s\t$%s\t\t%s\n",
				bct.WordFromInt(pos/2+origin).String(),
				t.String()[1:],
				instr.String(),
			)
		} else if instr.Lo != bct.Trybble(instruction.Nop) {
			switch instr.Lo {
			case instruction.LitTryte:
				t := bct.Tryte{Hi: bct.Trybble(content[pos+o]), Lo: bct.Trybble(content[pos+o+1])}
				fmt.Printf("%s\t$%s\t\t%s%s\n",
					bct.WordFromInt(pos/2+origin).String(),
					t.String(),
					instr.String(),
					t.String(),
				)
				o += 2
			case instruction.LitWord:
				w := bct.Word{
					Hi: bct.Tryte{Hi: bct.Trybble(content[pos+o]), Lo: bct.Trybble(content[pos+o+1])},
					Lo: bct.Tryte{Hi: bct.Trybble(content[pos+o+2]), Lo: bct.Trybble(content[pos+o+3])},
				}
				fmt.Printf("%s\t$%s\t\t%s%s\n",
					bct.WordFromInt(pos/2+origin).String(),
					w.String(),
					instr.String(),
					w.String(),
				)
				o += 4
			default:
				i := instruction.Simple(instr.Lo)
				fmt.Printf("%s\t%s\t\t%s\n", bct.WordFromInt(pos/2+origin).String(), i.String(), instr.String())
			}
		} else {
			i := instruction.Simple(instr.Hi)
			fmt.Printf("%s\t%s\t\t%s\n", bct.WordFromInt(pos/2+origin).String(), i.String(), instr.String())
		}

		pos += o
	}
}
