package sqlreflect

type Schema struct {
	CatalogName string `sql:"catalog_name"`
	SchemaName  string `sql:"schema_name"`
	SchemaOwner string `sql:"schema_owner"`

	DefaultCharacterSetCatalog string `sql:"default_character_set_catalog"`
	DefaultCharacterSetSchema  string `sql:"default_character_set_schema"`
	DefaultCharacterSetName    string `sql:"default_character_set_name"`

	SQLPath string `sql:"sql_path"`
}
