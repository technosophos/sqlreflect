package sqlreflect

import (
	"testing"

	"github.com/Masterminds/squirrel"
)

func loadTestTable(t *testing.T, name string) *Table {
	si := New(DBOptions{Driver: "postgres", Queryer: squirrel.NewStmtCacheProxy(db)})

	tt, err := si.Table(name, tCatalog, "")
	if err != nil {
		t.Fatal(err)
	}
	if tt.TableNameField != name {
		t.Fatalf("Expected %q , got %q", name, tt.TableNameField)
	}
	return tt
}

func TestTable_Privileges(t *testing.T) {
	table := loadTestTable(t, "person")
	// TODO: Actually call Privileges and check the output.
	privs, err := table.Privileges()
	if err != nil {
		t.Fatal(err)
	}
	if len(privs) == 0 {
		t.Fatalf("Expected at least one privilege on %s", table.TableNameField)
	}
}

func TestTable_Constraints(t *testing.T) {
	table := loadTestTable(t, "person")
	// TODO: Actually call Privileges and check the output.
	constraints, err := table.Constraints()
	if err != nil {
		t.Fatal(err)
	}
	if len(constraints) == 0 {
		t.Fatalf("Expected at least one privilege on %s", table.TableNameField)
	}
	for _, c := range constraints {
		t.Logf("CONSTRAINT: %v", c)
	}
}

func TestTable_ConstraintsByType(t *testing.T) {
	table := loadTestTable(t, "person")
	constraints, err := table.ConstraintsByType(ConstraintPrimaryKey)
	if err != nil {
		t.Fatal(err)
	}
	if len(constraints) != 1 {
		t.Fatalf("Expected one primary key on %s", table.TableNameField)
	}
	constraints, err = table.ConstraintsByType(ConstraintForeignKey)
	if err != nil {
		t.Fatal(err)
	}
	if len(constraints) != 0 {
		t.Fatalf("Expected 0 constraints, got %d", len(constraints))
	}
}

func TestTable_PrimaryKey(t *testing.T) {
	table := loadTestTable(t, "person")
	pk, err := table.PrimaryKey()
	if err != nil {
		t.Fatal(err)
	}
	if pk.ConstraintType != ConstraintPrimaryKey {
		t.Fatalf("Unexpected primary key constraint type: %q", pk.ConstraintType)
	}
}

func TestTable_ForeignKeys(t *testing.T) {
	table := loadTestTable(t, "person")
	fk, err := table.ForeignKeys()
	if err != nil {
		t.Fatal(err)
	}
	if len(fk) != 0 {
		t.Errorf("Expected no foreign keys for the person table, but got %d", len(fk))
	}

	table = loadTestTable(t, "employees")
	fk, err = table.ForeignKeys()
	if err != nil {
		t.Fatal(err)
	}
	if len(fk) != 2 {
		t.Fatalf("Expected 2 foreign keys for the employees table, but got %d", len(fk))
	}
}
