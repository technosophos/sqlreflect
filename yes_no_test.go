package sqlreflect

import (
	"testing"
)

func TestYesNo_Scan(t *testing.T) {
	y := new(YesNo)
	y.Scan("YES")

	if !y.Bool {
		t.Fatal("Expected scan of YES to set bool to true")
	}

	y.Scan("NO")
	if y.Bool {
		t.Fatal("Expected scan of NO to set bool to false")
	}
}

func TestYesNo_String(t *testing.T) {
	y := &YesNo{Bool: true}
	if y.String() != YesNoYes {
		t.Error("Expected bool true to produce YES")
	}

	y = &YesNo{}
	if y.String() != YesNoNo {
		t.Error("Expected bool false to produce NO")
	}
}
