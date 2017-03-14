package sqlreflect

import (
	"testing"
)

func TestTableConstraint_ColumnUsage(t *testing.T) {
	table := loadTestTable(t, "person")
	con, err := table.Constraint("person_pkey")
	if err != nil {
		t.Fatal(err)
	}

	usage, err := con.ColumnUsage()
	if err != nil {
		t.Fatal(err)
	}

	// There is only one column used for the primary key on org.
	if usage[0].ColumnName != "id" {
		t.Errorf("Expected column name id, got %q", usage[0].ColumnName)
	}
	if usage[0].OrdinalPosition != 1 {
		t.Errorf("Expected column ordinal pos to be 1, got %d", usage[0].OrdinalPosition)
	}
}
