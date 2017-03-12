package sqlreflect

import (
	"database/sql/driver"
	"fmt"
)

const (
	// YesNoYes is the SQL defined string for YES in a yes_or_no type.
	YesNoYes = "YES"
	// YesNoNo is the SQL-defined string for NO in a yes_or_no type.
	YesNoNo = "NO"
)

// YesNo takes an information schema's yes_or_no type and manages it as a bool.
type YesNo struct {
	// Bool is the boolean value of this flag.
	Bool bool
}

// Scan implements the database/sql.Scanner interface.
func (y *YesNo) Scan(v interface{}) error {
	if fmt.Sprintf("%v", v) == "YES" {
		y.Bool = true
	}
	return nil
}

// Value iplementes the database/sql/driver.Valuer interface.
func (y *YesNo) Value() (driver.Value, error) { return y.Bool, nil }

// String returns the YES/NO version of this flag.
//
// String implements fmt.Stringer.
func (y *YesNo) String() string {
	if y.Bool {
		return "YES"
	}
	return "NO"
}
