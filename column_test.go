package sqlreflect

import (
	"testing"
)

func TestColumn_Options(t *testing.T) {
	table := loadTestTable(t, "person")
	col, err := table.Column("id")
	if err != nil {
		t.Fatal(err)
	}
	opts, err := col.Options()
	if err != nil {
		t.Fatal(err)
	}
	if len(opts) != 0 {
		t.Errorf("Expected no options, got %d", len(opts))
	}
}
