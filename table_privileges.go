package sqlreflect

// TODO
type TablePrivilege struct {
	//TableLocator
	TableCatalog   string `stbl:"table_catalog"`
	TableSchema    string `stbl:"table_schema"`
	TableNameField string `stbl:"table_name"`
	// Core firelds
	Grantor       string `stbl:"grantor"`
	Grantee       string `stbl:"grantee"`
	PrivilegeType string `stbl:"privilege_type"`
	IsGrantable   YesNo  `stbl:"is_grantable"`
	WithHierarchy YesNo  `stbl:"with_hierarchy"`

	opts *DBOptions
}

func (t *TablePrivilege) TableName() string {
	return "information_schema.table_privileges"
}
