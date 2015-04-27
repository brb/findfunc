package findfunc

import "testing"

func TestFindFunc(t *testing.T) {
	cases := []struct {
		fileName string
		pc       uint64
		wantName string
		wantAddr uint64
		wantErr  error
	}{
		{"func", 0x40010c, "foo", 0x40010c, nil},
		{"func", 0x40010d, "foo", 0x40010c, nil},
		{"func", 0x400117, "bar", 0x400117, nil},
		{"func", 0x400000, "", 0, &ErrFunctionNotFound{}},
		{"nofunc", 0x400000, "", 0, &ErrNoFunctions{}},
	}

	for i, c := range cases {
		gotSymbol, gotErr := FindFunc("../testdata/"+c.fileName, c.pc)
		if gotSymbol.Name != c.wantName || gotSymbol.Value != c.wantAddr ||
			gotErr != c.wantErr {
			t.Errorf("test %d: expected (%s, %x, %s), got (%s, %x, %s)",
				i,
				c.wantName, c.wantAddr, c.wantErr,
				gotSymbol.Name, gotSymbol.Value, gotErr)

		}
	}
}
