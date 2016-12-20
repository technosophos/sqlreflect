package sqlreflect

// Column represents a column (attribute) attached to a table.
// A column can exist on exactly one table.
type Column struct {
	TableCatalog    string `sql:"table_catalog"`
	TableSchema     string `sql:"table_schema"`
	TableName       string `sql:"table_name"`
	Name            string `sql:"column_name"`
	OrdinalPosition int    `sql:"ordinal_position"`
	// Default is a description of the default value, not the actual default value.
	Default                string `sql:"column_default"`
	IsNullable             bool   `sql:"is_nullable"` // yes_or_no
	DataType               string `sql:"data_type"`
	CharacterMaximumLength int    `sql:"character_maximum_length"`
	CharacterOctetLength   int    `sql:"character_octet_length"`
	NumericPrecision       int    `sql:"numeric_precision"`
	NumericPrecisionRadix  int    `sql:"numeric_precision_radix"`
	NumericScale           int    `sql:"numeric_scale"`
	DatetimePrecision      int    `sql:"datetime_precision"`
	IntervalType           string `sql:"interval_type"`
	IntervalPrecision      int    `sql:"interval_precision"`
	CharacterSetCatalog    string `sql:"character_set_catalog"`
	CharacterSetSchema     string `sql:"character_set_schema"`
	CharacterSetName       string `sql:"character_set_name"`
	CollationCatalog       string `sql:"collation_catalog"`
	CollationSchema        string `sql:"collation_schema"`
	CollationName          string `sql:"collation_name"`
	DomainCatalog          string `sql:"domain_catalog"`
	DomainSchema           string `sql:"domain_schema"`
	DomainName             string `sql:"domain_name"`
	UDTCatalog             string `sql:"udt_catalog"`
	UDTSchema              string `sql:"udt_schema"`
	UDTName                string `sql:"udt_name"`
	ScopeCatalog           string `sql:"scope_catalog"`
	ScopeSchema            string `sql:"scope_schema"`
	ScopeName              string `sql:"scope_name"`
	MaximumCardinality     int    `sql:"maximum_cardinality"`
	IsSelfReferencing      bool   `sql:"is_self_referencing"`
	IsIdentity             bool   `sql:"is_identity"`
	IdentityGeneration     string `sql:"identity_generation"`
	IdentityStart          string `sql:"identity_start"`
	IdentityIncrement      string `sql:"identity_increment"`
	IdentityMaximum        string `sql:"identity_maximum"` // PG docs say string
	IdentityMinimum        string `sql:"identity_minimum"`
	IdentityCycle          bool   `sql:"identity_cycle"`
	IsGenerated            string `sql:"is_generated"` // PG docs say string
	GenerationExpression   string `sql:"generation_expression"`
	IsUpdatable            bool   `sql:"is_updatable"`
}

func (this *Column) Privileges() []*ColumnPrivilege {
	return []*ColumnPrivilege{}
}

func (this *Column) Options() []*ColumnOption {
	return []*ColumnOption{}
}

func (this *Column) DomainUsage() []*ColumnDomainUsage {
	return []*ColumnDomainUsage{}
}

func (this *Column) UDTUsage() []*ColumnUDTUsage {
	return []*ColumnUDTUsage{}
}

// Constrains returns a record of a constaint that references this column.
//
// For instance, if another table references this column as a foreign key,
// this will return information about that constraint.
func (this *Column) Constrains() []*ConstraintColumnUsage {
	return []*ConstraintColumnUsage{}
}

// Keys the restrictions placed on this column due to its use as a key.
//
// For example, if the column is a foreign key, this will return information
// about that foreign key relationship.
//
// This returns for primary key, foreign key, and uniqueness constraints.
func (this *Column) Keys() []*KeyColumnUsage {
	return []*KeyColumnUsage{}
}
