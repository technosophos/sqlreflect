package sqlreflect

type Table struct {
	//TableLocator
	TableCatalog   string `stbl:"table_catalog"`
	TableSchema    string `stbl:"table_schema"`
	TableNameField string `stbl:"table_name"`
	// TableType is normally one of BASE TABLE, VIEW, FOREIGN TABLE, or LOCAL TEMPORARY
	TableType                 string `stbl:"table_type"`
	SelfReferencingColumnName string `stbl:"self_referencing_column_name"`
	ReferenceGeneration       string `stbl:"reference_generation"`
	UserDefinedTypeCatalog    string `stbl:"user_defined_type_catalog"`
	UserDefinedTypeSchema     string `stbl:"user_defined_type_schema"`
	IsInsertableInto          bool   `stbl:"is_insertable_into"` // actual type is yes_no
	IsTyped                   bool   `stbl:"is_typed"`           // also yes_or_no
	CommitAction              string `stbl:"commit_action"`
}

func (t Table) TableName() string {
	return "information_schema.tables"
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
