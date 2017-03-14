package sqlreflect

type ColumnOption struct {
	TableCatalog   string `stbl:"table_catalog"`
	TableSchema    string `stbl:"table_schema"`
	TableNameField string `stbl:"table_name"`

	ColumnName  string `sql:"column_name"`
	OptionName  string `sql:"option_name"`
	OptionValue string `sql:"option_value"`

	opts *DBOptions
}

func (this ColumnOption) TableName() string {
	return "information_schema.column_options"
}
