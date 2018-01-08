package entity_test

import (
	"os"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/themccallister/entity"
)

func TestCanGetAnInMemoryDatabaseBasedOnAppEnv(t *testing.T) {
	os.Setenv("APP_ENV", "testing")

	sess, err := entity.NewSession()
	if err != nil {
		t.Fatalf("expected the error to not be nil, got `%v` instead", err)
	}

	err = sess.ORM.DB().Ping()
	if err != nil {
		t.Fatal("we should have been able to ping the db in memory")
	}

	os.Setenv("APP_ENV", "")
}
