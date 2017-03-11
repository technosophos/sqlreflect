package sqlreflect

import (
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
)

func optionalCatalogSchema(q squirrel.SelectBuilder, catalog, schema string) squirrel.SelectBuilder {
	if catalog != "" {
		q = q.Where("table_catalog = ?", catalog)
	}
	if schema == "" {
		// Is this postgres specific? If so, wrap in options check.
		schema = "public"
	}
	q = q.Where("table_schema = ?", schema)
	logQ(q)
	return q
}

func logQ(q squirrel.SelectBuilder) {
	a, b, c := q.ToSql()
	log.Printf(">>> %q // %s // %v", a, expandQP(b), c)
}

func expandQP(params []interface{}) string {
	p := "["
	for i, ii := range params {
		p += fmt.Sprintf(`$%d="%v" `, i+1, ii)
	}
	p += "]"
	return p
}
