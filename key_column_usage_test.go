package sqlreflect

import (
	"testing"
)

func TestKeyColumnUsage_Column(t *testing.T) {
	table := loadTestTable(t, "org")
	con, err := table.Constraint("org_pkey")
	if err != nil {
		t.Fatal(err)
	}

	usage, err := con.ColumnUsage()
	if err != nil {
		t.Fatal(err)
	}

	c, err := usage[0].Column()
	if err != nil {
		t.Fatal(err)
	}

	if c.Name != "id" {
		t.Errorf("Expected ordinal position 1 to be id, got %q", c.Name)
	}
	if c.IsNullable.Bool {
		t.Error("Expected id to not be nullable")
	}
	if c.DataType != "integer" {
		t.Errorf("Expected id DataType to be integer, got %q", c.DataType)
	}
}
