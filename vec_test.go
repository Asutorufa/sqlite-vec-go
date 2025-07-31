package sqlitevecgo

import (
	"fmt"
	"testing"

	"github.com/ncruces/go-sqlite3/driver"
)

func TestVec(t *testing.T) {
	db, err := driver.Open(":memory:")
	if err != nil {
		panic(err)
	}

	var sqlite_version, vec_version string
	err = db.QueryRow(`SELECT sqlite_version(), vec_version()`).Scan(&sqlite_version, &vec_version)
	if err != nil {
		panic(err)
	}

	fmt.Printf("sqlite_version=%s, vec_version=%s\n", sqlite_version, vec_version)
}
