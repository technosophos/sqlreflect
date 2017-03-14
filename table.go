package sqlreflect

import (
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/Masterminds/structable"
)

type Table struct {
	//TableLocator
	TableCatalog   string `stbl:"table_catalog"`
	TableSchema    string `stbl:"table_schema"`
	TableNameField string `stbl:"table_name"`
	// TableType is normally one of BASE TABLE, VIEW, FOREIGN TABLE, or LOCAL TEMPORARY
	TableType                 string `stbl:"table_type"`
	SelfReferencingColumnName string `stbl:"self_referencing_column_name"`
	ReferenceGeneration       string `stbl:"reference_generation"`
	UserDefinedTypeCatalog    string `stbl:"user_defined_type_catalog"`
	UserDefinedTypeSchema     string `stbl:"user_defined_type_schema"`
	IsInsertableInto          bool   `stbl:"is_insertable_into"` // actual type is yes_no
	IsTyped                   bool   `stbl:"is_typed"`           // also yes_or_no
	CommitAction              string `stbl:"commit_action"`
	opts                      *DBOptions
}

func (t Table) TableName() string {
	return "information_schema.tables"
}

// Privileges returns the table privileges for this table.
func (this *Table) Privileges() ([]*TablePrivilege, error) {
	t := &TablePrivilege{}
	res := []*TablePrivilege{}
	st := structable.New(this.opts.Queryer, this.opts.Driver).Bind(t.TableName(), t)
	fn := func(d structable.Describer, q squirrel.SelectBuilder) (squirrel.SelectBuilder, error) {
		q = q.Where("table_schema=? AND table_catalog = ? AND table_name = ?",
			this.TableSchema, this.TableCatalog, this.TableNameField)
		return q, nil
	}
	items, err := structable.ListWhere(st, fn)
	if err != nil {
		return res, err
	}
	for _, item := range items {
		tt := item.Interface().(*TablePrivilege)
		tt.opts = this.opts
		res = append(res, tt)
	}

	return res, nil
}

// Constraints returns the constraints imposed on this table.
func (this *Table) Constraints() ([]*TableConstraint, error) {
	t := &TableConstraint{}
	res := []*TableConstraint{}
	st := structable.New(this.opts.Queryer, this.opts.Driver).Bind(t.TableName(), t)
	fn := func(d structable.Describer, q squirrel.SelectBuilder) (squirrel.SelectBuilder, error) {
		q = q.Where("table_schema=? AND table_catalog = ? AND table_name = ?",
			this.TableSchema, this.TableCatalog, this.TableNameField)
		return q, nil
	}
	items, err := structable.ListWhere(st, fn)
	if err != nil {
		return res, err
	}
	for _, item := range items {
		tt := item.Interface().(*TableConstraint)
		tt.opts = this.opts
		res = append(res, tt)
	}
	return res, nil
}

func (this *Table) ConstraintsByType(name ConstraintType) ([]*TableConstraint, error) {
	t := &TableConstraint{}
	res := []*TableConstraint{}
	st := structable.New(this.opts.Queryer, this.opts.Driver).Bind(t.TableName(), t)
	fn := func(d structable.Describer, q squirrel.SelectBuilder) (squirrel.SelectBuilder, error) {
		q = q.Where("table_schema=? AND table_catalog = ? AND table_name = ? AND constraint_type = ?",
			this.TableSchema, this.TableCatalog, this.TableNameField, string(name))
		return q, nil
	}
	items, err := structable.ListWhere(st, fn)
	if err != nil {
		return res, err
	}
	for _, item := range items {
		tt := item.Interface().(*TableConstraint)
		tt.opts = this.opts
		res = append(res, tt)
	}
	return res, nil
}

// Columns returns the columns contained by this table.
func (this *Table) Columns() ([]*Column, error) {
	t := &Column{}
	res := []*Column{}
	st := structable.New(this.opts.Queryer, this.opts.Driver).Bind(t.TableName(), t)
	fn := func(d structable.Describer, q squirrel.SelectBuilder) (squirrel.SelectBuilder, error) {
		q = q.Where("table_schema=? AND table_catalog = ? AND table_name = ?",
			this.TableSchema, this.TableCatalog, this.TableNameField)
		return q, nil
	}
	items, err := structable.ListWhere(st, fn)
	if err != nil {
		return res, err
	}
	for _, item := range items {
		tt := item.Interface().(*Column)
		tt.opts = this.opts
		res = append(res, tt)
	}
	return res, nil
}

// PrimaryKey returns the primary key for this table.
//
// TODO: Is it ever possible to have two table constraints for one primary
// key?
func (this *Table) PrimaryKey() (*TableConstraint, error) {
	tt, err := this.ConstraintsByType(ConstraintPrimaryKey)
	if err != nil {
		return nil, err
	}
	if len(tt) == 0 {
		return nil, errors.New("No primary key")
	}
	return tt[0], nil
}

// ForeignKeys returns a list of foreign key table constraints.
func (this *Table) ForeignKeys() ([]*TableConstraint, error) {
	return this.ConstraintsByType(ConstraintForeignKey)
}

// InViews returns a list of views that use this table.
//
// Not that this might not be the only table used by that view.
func (this *Table) InViews() ([]*View, error) {
	q := squirrel.Select("view_name", "view_catalog", "view_schema").
		From("information_schema.view_table_usage").
		Where(`table_catalog = ? AND table_schema = ? AND table_name = ?`,
			this.TableCatalog, this.TableSchema, this.TableNameField).
		RunWith(this.opts.Queryer)

	if this.opts.Driver == "postgres" {
		q = q.PlaceholderFormat(squirrel.Dollar)
	}

	rows, err := q.Query()
	if err != nil {
		return []*View{}, err
	}
	defer rows.Close()

	vs := []*View{}
	for rows.Next() {
		v := &View{}
		rows.Scan(&v.TableNameField, &v.TableCatalog, &v.TableSchema)
		st := structable.New(this.opts.Queryer, this.opts.Driver).Bind(v.TableName(), v)
		if err := st.LoadWhere("table_catalog = ? AND table_schema = ? AND table_name = ?",
			v.TableCatalog, v.TableSchema, v.TableNameField); err != nil {
			return vs, err
		}
		vs = append(vs, v)
	}
	return vs, rows.Err()
}
