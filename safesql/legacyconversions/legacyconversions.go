package legacyconversions

import (
	sql "safesql"
	"safesql/internal/raw"
)

var trustedSQLCtor = raw.TrustedSQLCtor.(func(string) sql.TrustedSQL)

func RiskilyAssumeTrustedSQL(trusted string) sql.TrustedSQL {
	return trustedSQLCtor(trusted)
}
