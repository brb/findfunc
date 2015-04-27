/*
Package findfunc implements a routine to find a function in
statically linked ELF executable by Program Counter (PC)
*/
package findfunc

import (
	"debug/elf"
	"sort"
)

type ErrFunctionNotFound struct{}
type ErrNoFunctions struct{}

func (e *ErrFunctionNotFound) Error() string {
	return "Function not found"
}

func (e *ErrNoFunctions) Error() string {
	return "Object file does not contain any function"
}

type Symbols []elf.Symbol

func (s Symbols) Len() int {
	return len(s)
}

func (s Symbols) Less(i, j int) bool {
	return s[i].Value < s[j].Value
}

func (s Symbols) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func FindFunc(fileName string, pc uint64) (s elf.Symbol, err error) {
	file, err := elf.Open(fileName)
	if err != nil {
		return elf.Symbol{}, err
	}

	symbols, err := file.Symbols()
	if err != nil {
		return elf.Symbol{}, err
	}

	funcSymbols := make(Symbols, 0)
	for _, symbol := range symbols {
		if elf.ST_TYPE(symbol.Info) == elf.STT_FUNC {
			funcSymbols = append(funcSymbols, symbol)
		}
	}
	n := len(funcSymbols)
	if n == 0 {
		return elf.Symbol{}, &ErrNoFunctions{}
	}

	sort.Sort(funcSymbols)
	i := sort.Search(n, func(i int) bool { return funcSymbols[i].Value >= pc })
	if i != 0 || funcSymbols[i].Value <= pc {
		if i == n || (funcSymbols[i].Value != pc && i != 0) {
			i--
		}
		return funcSymbols[i], nil
	}

	return elf.Symbol{}, &ErrFunctionNotFound{}
}
