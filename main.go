package main

import (
	"flag"
	"fmt"
	"github.com/brb/findfunc/findfunc"
	"os"
	"strconv"
)

func main() {
	fileName := flag.String("f", "", "Path to ELF executable")
	pc := flag.String("pc", "", "Program Counter")
	isHex := flag.Bool("x", false, "Is PC in hex?")
	flag.Parse()

	if *fileName == "" || *pc == "" {
		flag.Usage()
	}

	base := 10
	if *isHex {
		base = 16
	}
	pcVal, err := strconv.ParseUint(*pc, base, 64)

	sym, err := findfunc.FindFunc(*fileName, pcVal)
	if err == nil {
		fmt.Printf("%s %x\n", sym.Name, sym.Value)
	} else {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(-1)
	}
}
