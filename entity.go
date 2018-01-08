package entity

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Session is a struct that contains a pointer to a GORM database. To create a new session, without passing one around, we rely on envrionment variables.
// Use the NewSession func to creatings a new session.
type Session struct {
	ORM *gorm.DB
}

// NewSession uses envrionment variables to create a new services to connect to a GORM database.
// if the APP_ENV is set to testing, an in-memory database (sqlite) will be used for easy testing.
func NewSession() (*Session, error) {
	switch os.Getenv("APP_ENV") {
	case "testing":
		db, err := gorm.Open("sqlite3", ":memory")
		if err != nil {
			return nil, err
		}
		return &Session{ORM: db}, nil
	default:
		db, err := gorm.Open("postgres", fmt.Sprintf(
			"host=%v user=%v dbname=%v sslmode=%v password=%v",
			os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_MODE"), os.Getenv("DB_PASS"),
		))
		if err != nil {
			return nil, err
		}
		return &Session{ORM: db}, nil
	}
}
