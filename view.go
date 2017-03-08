package sqlreflect

type View struct {
	//TableLocator
	TableCatalog   string `stbl:"table_catalog"`
	TableSchema    string `stbl:"table_schema"`
	TableNameField string `stbl:"table_name"`
	// View definition
	ViewDefinition          string `sql:"view_definition"`
	CheckOption             string `sql:"check_option"`
	IsUpdatable             bool   `sql:"is_updatable"`
	IsInsertableInto        bool   `sql:"is_insertable_into"`
	IsTriggerUpdatable      bool   `sql:"is_trigger_updatable"`
	IsTriggerDeletable      bool   `sql:"is_trigger_deletable"`
	IsTriggerInsertableInto bool   `sql:"is_trigger_insertable_into"`
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
	ViewCatalog string `sql:"view_catalog"`
	ViewSchema  string `sql:"view_schema"`
	ViewName    string `sql:"view_name"`
}
