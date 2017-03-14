package sqlreflect

// TODO
type TableConstraint struct {
	//TableLocator
	TableCatalog   string `stbl:"table_catalog"`
	TableSchema    string `stbl:"table_schema"`
	TableNameField string `stbl:"table_name"`

	// ConstraintLocator
	ConstraintCatalog string `stbl:"constraint_catalog"`
	ConstraintSchema  string `stbl:"constraint_schema"`
	ConstraintName    string `stbl:"constraint_name"`

	ConstraintType    ConstraintType `stbl:"constraint_type"`
	IsDeferrable      YesNo          `stbl:"is_deferrable"`
	InitiallyDeferred YesNo          `stbl:"initially_deferred"`

	opts *DBOptions
}

func (t *TableConstraint) TableName() string {
	return "information_schema.table_constraints"
}
