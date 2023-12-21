package db

import (
	"database/sql"
	"log"

	"github.com/leonkay/code_sandbox/golang/bragdoc/config"
)

type DbHandle interface {
	Db() *sql.DB
	Version() string
	Create()
}

type DbHandleFactory interface {
	Connect(config.Config) DbHandle
}

var ConnectionFactoryMap = map[string]DbHandleFactory{
	"sqlite3": &SqliteHandler{},
}

func Run(config config.Config) DbHandle {
	factory := ConnectionFactoryMap["sqlite3"]
	handler := factory.Connect(config)
	handler.Create()
	return handler
}

func Close(db DbHandle) {
	err := db.Db().Close()
	if err != nil {
		log.Fatal("Error Closing Database Connection")
	}
}
