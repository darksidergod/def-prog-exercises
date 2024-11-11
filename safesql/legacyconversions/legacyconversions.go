package legacyconversions

import (
	sql "github.com/darksidergod/def-prog-exercises/safesql"
	"github.com/darksidergod/def-prog-exercises/safesql/internal/raw"
)

var trustedSQLCtor = raw.TrustedSQLCtor.(func(string) sql.TrustedSQL)

func RiskilyAssumeTrustedSQL(trusted string) sql.TrustedSQL {
	return trustedSQLCtor(trusted)
}
