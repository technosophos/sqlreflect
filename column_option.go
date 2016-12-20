package sqlreflect

type ColumnOption struct {
	TableLocator
	ColumnName  string `sql:"column_name"`
	OptionName  string `sql:"option_name"`
	OptionValue string `sql:"option_value"`
}
