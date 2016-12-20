package sqlreflect

type Table struct {
	TableLocator
	// TableType is normally one of BASE TABLE, VIEW, FOREIGN TABLE, or LOCAL TEMPORARY
	TableType                 string `sql:"table_type"`
	SelfReferencingColumnName string `sql:"self_referencing_column_name"`
	ReferenceGeneration       string `sql:"reference_generation"`
	UserDefinedTypeCatalog    string `sql:"user_defined_type_catalog"`
	UserDefinedTypeSchema     string `sql:"user_defined_type_schema"`
	IsInsertableInto          bool   `sql:"is_insertable_into"` // actual type is yes_no
	IsTyped                   bool   `sql:"is_typed"`           // also yes_or_no
	CommitAction              string `sql:"commit_action"`
}

// Privileges returns the table privileges for this table.
func (this *Table) Privileges() []*TablePrivilege {
	return []*TablePrivilege{}
}

// Constraints returns the constraints imposed on this table.
func (this *Table) Constraints() []*TableConstraint {
	return []*TableConstraint{}
}

// Columns returns the columns contained by this table.
func (this *Table) Columns() []*Column {
	return []*Column{}
}

// PrimaryKey returns the primary key for this table.
//
// TODO: Is it ever possible to have two table constraints for one primary
// key.
func (this *Table) PrimaryKey() *TableConstraint {
	return &TableConstraint{}
}

// ForeignKeys returns a list of foreign key table constraints.
func (this *Table) ForeignKeys() []*TableConstraint {
	return []*TableConstraint{}
}

// InViews returns a list of views that use this table.
//
// Not that this might not be the only table used by that view.
func (this *Table) InViews() []*View {
	return []*View{}
}
