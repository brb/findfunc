# Find Function by Program Counter (PC) in ELF Executable

As the title states, the program is used to find a function name when only PC is
known. Useful when debugging segmentation faults, unhandled page-faults, etc of
big binaries.

## Assumptions

* Statically linked ELF executable.
* No relocations by a loader.
* Absolute value of PC is given.

## Usage example

```bash
$ go install github.com/brb/findfunc
$ $GOPATH/bin/findfunc -f $GOPATH/src/github.com/brb/findfunc/testdata/func -pc 400121 -x
bar 400117
```
