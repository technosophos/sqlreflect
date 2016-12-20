package sqlreflect

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
)

// Queryer defines an interface for querying databases
type Queryer interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

// Preparer provides support for prepared statements.
type Preparer interface {
	Prepare(query string) (*sql.Stmt, error)
}

// DBOptions describes the database queries are to be executed against.
type DBOptions struct {
	// Placeholder describes the placeholder format. If left nil, this will
	// use '?' as the placeholder. Postgres users may prefer to set this to
	// *squirrel.Dollar, which uses $1, $2... instead of ?, ?...
	Placeholder squirrel.PlaceholderFormat

	// Querier is the runner that can execute database queries.
	// Often, this is just a *sql.DB. If the given queryer also implements
	// Preparer, this will use cached prepared statements instead of
	// directly executed queries. To disable the use of a query cache,
	// set DisableCache to true.
	Queryer Queryer

	// DisableCache disables using a query cache and prepared statements.
	DisableCache bool
}

// SchemaInfo provides access to the database schemata.
type SchemaInfo struct {
	opts   *DBOptions
	runner squirrel.Queryer
}

// New creates a new SchemaInfo.
func New(opts DBOptions) *SchemaInfo {
	if opts.Placeholder == nil {
		opts.Placeholder = squirrel.Question
	}

	s := &SchemaInfo{
		opts:   &opts,
		runner: opts.Queryer,
	}

	_, isPrep := opts.Queryer.(Preparer)
	if isPrep && !opts.DisableCache {
		s.runner = squirrel.NewStmtCacher(opts.Queryer.(Preparer))
	} else {
		s.runner = opts.Queryer
	}

	return s
}

// Supported returns true if the given database supports Schema Info.
//
// When running on an unknown database or driver, this can be used as a
// general mechanism to report whether any of the functions in this library
// can return useful results.
func (s *SchemaInfo) Supported() bool {
	return false
}

func (s *SchemaInfo) Tables() []*Table {
	return []*Table{}
}
