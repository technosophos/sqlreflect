package sqlreflect

import (
	"database/sql"
	"database/sql/driver"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/Masterminds/structable"
)

// Queryer defines an interface for querying databases
type Queryer interface {
	squirrel.DBProxyBeginner
	//Query(query string, args ...interface{}) (*sql.Rows, error)
	//QueryRow(query string, args ...interface{}) *sql.Row

	//// We need exec because squirrel's select builder requires it.

	//Exec(query string, args ...interface{}) (sql.Result, error)
}

// Preparer provides support for prepared statements.
type Preparer interface {
	Prepare(query string) (*sql.Stmt, error)
}

// DBOptions describes the database queries are to be executed against.
type DBOptions struct {
	// Driver is the string name of the registered driver.
	Driver string

	// Querier is the runner that can execute database queries.
	// Often, this is just a *sql.DB. If the given queryer also implements
	// Preparer, this will use cached prepared statements instead of
	// directly executed queries. To disable the use of a query cache,
	// set DisableCache to true.
	//
	// FIXME: Right now, this is a squirrel.StmtCacheProxy
	Queryer Queryer

	// DisableCache disables using a query cache and prepared statements.
	DisableCache bool
}

func NewDBOptions(db *sql.DB, driver string) DBOptions {
	return DBOptions{Driver: driver, Queryer: squirrel.NewStmtCacheProxy(db)}
}

// DbRecorder returns a new structable.DbRecorder ready to be bound to a table.
func (d *DBOptions) DbRecorder() *structable.DbRecorder {
	return structable.New(d.Queryer, d.Driver)
}

// SchemaInfo provides access to the database schemata.
type SchemaInfo struct {
	opts   *DBOptions
	runner squirrel.BaseRunner

	// placeholder describes the placeholder format. If left nil, this will
	// use '?' as the placeholder. Postgres users may prefer to set this to
	// *squirrel.Dollar, which uses $1, $2... instead of ?, ?...
	placeholder squirrel.PlaceholderFormat
}

// New creates a new SchemaInfo.
func New(opts DBOptions) *SchemaInfo {
	s := &SchemaInfo{
		opts:   &opts,
		runner: opts.Queryer,
	}

	s.placeholder = squirrel.Question
	if opts.Driver == "postgres" {
		s.placeholder = squirrel.Dollar
	}

	/* Once preparer and queryer are fixed, uncomment this.
	_, isPrep := opts.Queryer.(Preparer)
	if isPrep && !opts.DisableCache {
		s.runner = squirrel.NewStmtCacher(opts.Queryer.(Preparer))
	} else {
		s.runner = opts.Queryer
	}
	*/
	s.runner = opts.Queryer

	return s
}

// Supported returns true if the given database supports Schema Info.
//
// When running on an unknown database or driver, this can be used as a
// general mechanism to report whether any of the functions in this library
// can return useful results.
func (s *SchemaInfo) Supported() bool {
	res, err := s.Select("catalog_name").
		From("information_schema.information_schema_catalog_name").
		Query()
	defer res.Close()
	return err == nil
}

// Tables gets all tables.
// Setting catalog and schema to something other than an empty string will
// constrain by those.
func (s *SchemaInfo) Tables(catalog, schema string) ([]*Table, error) {
	t := &Table{}
	st := structable.New(s.opts.Queryer, s.opts.Driver).Bind(t.TableName(), t)
	fn := func(d structable.Describer, q squirrel.SelectBuilder) (squirrel.SelectBuilder, error) {
		q = q.Where(`table_type='BASE TABLE'`)
		return optionalCatalogSchema(q, catalog, schema), nil
	}
	items, err := structable.ListWhere(st, fn)
	if err != nil {
		return []*Table{}, err
	}
	tables := make([]*Table, len(items))
	for i, item := range items {
		tt := item.Interface().(*Table)
		tt.opts = s.opts
		tables[i] = tt
	}
	return tables, nil
}

// Table gets a table by name (required).
// If catalog and schema are set, those will be used as additional constraints.
// Constraining by catalog (e.g. database) is highly recommended.
func (s *SchemaInfo) Table(name, catalog, schema string) (*Table, error) {
	t := &Table{}
	st := structable.New(s.opts.Queryer, s.opts.Driver).Bind(t.TableName(), t)
	fn := func(d structable.Describer, q squirrel.SelectBuilder) (squirrel.SelectBuilder, error) {
		// Basically, remove the default limits.
		q = q.Where(`table_type='BASE TABLE' AND table_name = ?`, name)
		return optionalCatalogSchema(q, catalog, schema), nil
	}
	items, err := structable.ListWhere(st, fn)
	if err != nil {
		return t, err
	}

	if l := len(items); l != 1 {
		return t, fmt.Errorf("Expected 1 table, got %d", l)
	}
	table := items[0].Interface().(*Table)
	table.opts = s.opts
	return table, nil
}

func (s *SchemaInfo) Views(catalog, schema string) ([]*View, error) {
	view := &View{}
	st := structable.New(s.opts.Queryer, s.opts.Driver).Bind(view.TableName(), view)
	fn := func(d structable.Describer, q squirrel.SelectBuilder) (squirrel.SelectBuilder, error) {
		return optionalCatalogSchema(q, catalog, schema), nil
	}
	items, err := structable.ListWhere(st, fn)
	if err != nil {
		return []*View{}, err
	}
	views := make([]*View, len(items))
	for i, item := range items {
		vv := item.Interface().(*View)
		vv.opts = s.opts
		views[i] = vv
	}
	return views, nil
}

func (s *SchemaInfo) View(name, catalog, schema string) (*View, error) {
	view := &View{}
	st := structable.New(s.opts.Queryer, s.opts.Driver).Bind(view.TableName(), view)
	fn := func(d structable.Describer, q squirrel.SelectBuilder) (squirrel.SelectBuilder, error) {
		// View names are stored in table_name, since a view is a table.
		q = q.Where("table_name=?", name)
		return optionalCatalogSchema(q, catalog, schema), nil
	}
	items, err := structable.ListWhere(st, fn)
	if err != nil {
		return &View{}, err
	}
	if l := len(items); l != 1 {
		return view, fmt.Errorf("Expected 1 view, got %d", l)
	}
	vv := items[0].Interface().(*View)
	vv.opts = s.opts
	return vv, nil
}

func (s *SchemaInfo) Select(columns ...string) squirrel.SelectBuilder {
	return squirrel.Select(columns...).RunWith(s.runner).PlaceholderFormat(s.placeholder)
}

type YesNo bool

func (y YesNo) Scan(v interface{}) error {
	if fmt.Sprintf("%v", v) == "YES" {
		y = true
	}
	return nil
}
func (y YesNo) Value() (driver.Value, error) { return y, nil }
