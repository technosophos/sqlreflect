package sqlreflect

type View struct {
	//TableLocator
	TableCatalog   string `stbl:"table_catalog"`
	TableSchema    string `stbl:"table_schema"`
	TableNameField string `stbl:"table_name"`
	// View definition
	ViewDefinition          string `stbl:"view_definition"`
	CheckOption             string `stbl:"check_option"`
	IsUpdatable             YesNo  `stbl:"is_updatable"`
	IsInsertableInto        YesNo  `stbl:"is_insertable_into"`
	IsTriggerUpdatable      YesNo  `stbl:"is_trigger_updatable"`
	IsTriggerDeletable      YesNo  `stbl:"is_trigger_deletable"`
	IsTriggerInsertableInto YesNo  `stbl:"is_trigger_insertable_into"`

	opts *DBOptions
}

func (v *View) Tables() []*Table {
	return []*Table{}
}

func (v *View) TableName() string {
	return "information_schema.views"
}

// Columns returns a list of columns for this view.
//
// Columns may be from different tables.
func (v *View) Columns() []*Column {
	return []*Column{}
}

type ViewLocator struct {
	ViewCatalog string `stbl:"view_catalog"`
	ViewSchema  string `stbl:"view_schema"`
	ViewName    string `stbl:"view_name"`
}
