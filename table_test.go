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
