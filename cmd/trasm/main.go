package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/patrislav/triton-tools/assembly"
	"github.com/patrislav/triton-tools/assembly/lexer"
)

func main() {
	// inputName := "./assembly/examples/display.tras"
	// outputName := "./display.bct"

	content, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Could not read input file: %s", err)
	}

	lex := lexer.New(string(content))
	tokens := make([]lexer.Token, 0)
	for tok := lex.NextToken(); tok.Kind != lexer.TokenEOF; tok = lex.NextToken() {
		tokens = append(tokens, tok)
	}

	origin := 9842
	prs := assembly.NewParser(tokens)
	prog := prs.ParseProgram(origin)
	if err := prog.ResolveAddresses(); err != nil {
		log.Fatalf("Could not resolve addresses: %s", err)
	}

	buf := prog.GetBytes()
	if _, err := os.Stdout.Write(buf); err != nil {
		log.Fatalf("Could not write program to file: %s", err)
	}
}
