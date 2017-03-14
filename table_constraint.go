package sqlreflect

import (
	"github.com/Masterminds/squirrel"
	"github.com/Masterminds/structable"
)

// TableConstraint defines a constraint (e.g. primary key, foreign key...)
// placed on a table.
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

// ColumnUsage returns information about which columns are used in this key.
func (this *TableConstraint) ColumnUsage() ([]*KeyColumnUsage, error) {
	t := &KeyColumnUsage{}
	res := []*KeyColumnUsage{}
	st := structable.New(this.opts.Queryer, this.opts.Driver).Bind(t.TableName(), t)
	fn := func(d structable.Describer, q squirrel.SelectBuilder) (squirrel.SelectBuilder, error) {
		q = q.Where("constraint_schema=? AND constraint_catalog = ? AND constraint_name = ?",
			this.ConstraintSchema, this.ConstraintCatalog, this.ConstraintName)
		return q, nil
	}
	items, err := structable.ListWhere(st, fn)
	if err != nil {
		return res, err
	}
	for _, item := range items {
		tt := item.Interface().(*KeyColumnUsage)
		tt.opts = this.opts
		res = append(res, tt)
	}
	return res, nil
}
