package sqlreflect

import (
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/Masterminds/structable"
)

type KeyColumnUsage struct {
	// ConstraintLocator
	ConstraintCatalog string `stbl:"constraint_catalog"`
	ConstraintSchema  string `stbl:"constraint_schema"`
	ConstraintName    string `stbl:"constraint_name"`

	//TableLocator
	TableCatalog   string `stbl:"table_catalog"`
	TableSchema    string `stbl:"table_schema"`
	TableNameField string `stbl:"table_name"`

	ColumnName                 string `stbl:"column_name"`
	OrdinalPosition            int    `stbl:"ordinal_position"`
	PositionInUniqueConstraint int    `stbl:"position_in_unique_constraint"`

	opts *DBOptions
}

func (this KeyColumnUsage) TableName() string {
	return "information_schema.key_column_usage"
}

// Column gets the column associated with this usage record.
func (this KeyColumnUsage) Column() (*Column, error) {
	t := &Column{}
	st := structable.New(this.opts.Queryer, this.opts.Driver).Bind(t.TableName(), t)
	fn := func(d structable.Describer, q squirrel.SelectBuilder) (squirrel.SelectBuilder, error) {
		q = q.Where("table_schema=? AND table_catalog = ? AND table_name = ? AND column_name = ?",
			this.TableSchema, this.TableCatalog, this.TableNameField, this.ColumnName).Limit(1)
		return q, nil
	}
	items, err := structable.ListWhere(st, fn)
	if err != nil {
		return t, err
	}

	if len(items) != 1 {
		return t, errors.New("named column not found")
	}

	item := items[0].Interface().(*Column)
	item.opts = this.opts
	return item, nil
}
