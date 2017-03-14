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

func TestColumn_Constraints(t *testing.T) {
	table := loadTestTable(t, "person")
	col, err := table.Column("id")
	if err != nil {
		t.Fatal(err)
	}

	con, err := col.Constraints()
	if err != nil {
		t.Fatal(err)
	}

	// We expect one primary key and two foreign key references.
	if l := len(con); l != 3 {
		t.Errorf("Expected 3 constraints, got %d", l)
		for _, ll := range con {
			t.Logf("Con: %v", ll)
		}
	}
}
