package sqlreflect

import (
	"github.com/Masterminds/squirrel"
	"github.com/Masterminds/structable"
)

// Column represents a column (attribute) attached to a table.
// A column can exist on exactly one table.
type Column struct {
	TableCatalog   string `stbl:"table_catalog"`
	TableSchema    string `stbl:"table_schema"`
	TableNameField string `stbl:"table_name"`

	Name            string `stbl:"column_name"`
	OrdinalPosition int    `stbl:"ordinal_position"`
	// Default is a description of the default value, not the actual default value.
	Default                string `stbl:"column_default"`
	IsNullable             YesNo  `stbl:"is_nullable"` // yes_or_no
	DataType               string `stbl:"data_type"`
	CharacterMaximumLength int    `stbl:"character_maximum_length"`
	CharacterOctetLength   int    `stbl:"character_octet_length"`
	NumericPrecision       int    `stbl:"numeric_precision"`
	NumericPrecisionRadix  int    `stbl:"numeric_precision_radix"`
	NumericScale           int    `stbl:"numeric_scale"`
	DatetimePrecision      int    `stbl:"datetime_precision"`
	IntervalType           string `stbl:"interval_type"`
	IntervalPrecision      int    `stbl:"interval_precision"`
	CharacterSetCatalog    string `stbl:"character_set_catalog"`
	CharacterSetSchema     string `stbl:"character_set_schema"`
	CharacterSetName       string `stbl:"character_set_name"`
	CollationCatalog       string `stbl:"collation_catalog"`
	CollationSchema        string `stbl:"collation_schema"`
	CollationName          string `stbl:"collation_name"`
	DomainCatalog          string `stbl:"domain_catalog"`
	DomainSchema           string `stbl:"domain_schema"`
	DomainName             string `stbl:"domain_name"`
	UDTCatalog             string `stbl:"udt_catalog"`
	UDTSchema              string `stbl:"udt_schema"`
	UDTName                string `stbl:"udt_name"`
	ScopeCatalog           string `stbl:"scope_catalog"`
	ScopeSchema            string `stbl:"scope_schema"`
	ScopeName              string `stbl:"scope_name"`
	MaximumCardinality     int    `stbl:"maximum_cardinality"`
	IsSelfReferencing      YesNo  `stbl:"is_self_referencing"`
	IsIdentity             YesNo  `stbl:"is_identity"`
	IdentityGeneration     string `stbl:"identity_generation"`
	IdentityStart          string `stbl:"identity_start"`
	IdentityIncrement      string `stbl:"identity_increment"`
	IdentityMaximum        string `stbl:"identity_maximum"` // PG docs say string
	IdentityMinimum        string `stbl:"identity_minimum"`
	IdentityCycle          YesNo  `stbl:"identity_cycle"`
	IsGenerated            string `stbl:"is_generated"` // PG docs say string
	GenerationExpression   string `stbl:"generation_expression"`
	IsUpdatable            YesNo  `stbl:"is_updatable"`

	opts *DBOptions
}

func (this Column) TableName() string {
	return "information_schema.columns"
}

func (this *Column) Privileges() []*ColumnPrivilege {
	return []*ColumnPrivilege{}
}

// Options returns a list of column options set for this column.
func (this *Column) Options() ([]*ColumnOption, error) {
	t := &ColumnOption{}
	res := []*ColumnOption{}
	st := structable.New(this.opts.Queryer, this.opts.Driver).Bind(t.TableName(), t)
	fn := func(d structable.Describer, q squirrel.SelectBuilder) (squirrel.SelectBuilder, error) {
		q = q.Where("table_schema=? AND table_catalog = ? AND table_name = ? AND column_name = ?",
			this.TableSchema, this.TableCatalog, this.TableNameField, this.Name)
		return q, nil
	}
	items, err := structable.ListWhere(st, fn)
	if err != nil {
		return res, err
	}
	for _, item := range items {
		tt := item.Interface().(*ColumnOption)
		tt.opts = this.opts
		res = append(res, tt)
	}

	return res, nil
}

func (this *Column) DomainUsage() []*ColumnDomainUsage {
	return []*ColumnDomainUsage{}
}

func (this *Column) UDTUsage() []*ColumnUDTUsage {
	return []*ColumnUDTUsage{}
}

// Constrains returns a record of constraints that reference this column.
//
// For instance, if another table references this column as a foreign key,
// this will return information about that constraint. It will also include
// constraints on this table that reference this column (e.g. primary key).
func (this *Column) Constraints() ([]*TableConstraint, error) {
	q := squirrel.Select("constraint_name", "constraint_catalog", "constraint_schema").
		From("information_schema.constraint_column_usage").
		Where(`table_catalog = ? AND table_schema = ? AND table_name = ? AND column_name = ?`,
			this.TableCatalog, this.TableSchema, this.TableNameField, this.Name).
		RunWith(this.opts.Queryer)

	if this.opts.Driver == "postgres" {
		q = q.PlaceholderFormat(squirrel.Dollar)
	}

	logQ(q)

	rows, err := q.Query()
	if err != nil {
		return []*TableConstraint{}, err
	}
	defer rows.Close()

	constraints := []*TableConstraint{}
	for rows.Next() {
		c := &TableConstraint{}
		rows.Scan(&c.ConstraintName, &c.ConstraintCatalog, &c.ConstraintSchema)
		st := structable.New(this.opts.Queryer, this.opts.Driver).Bind(c.TableName(), c)
		if err := st.LoadWhere("constraint_catalog = ? AND constraint_schema = ? AND constraint_name = ?",
			c.ConstraintCatalog, c.ConstraintSchema, c.ConstraintName); err != nil {
			return constraints, err
		}
		constraints = append(constraints, c)
	}
	return constraints, rows.Err()
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
