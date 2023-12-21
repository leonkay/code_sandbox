package db

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/leonkay/code_sandbox/golang/bragdoc/config"
	"github.com/leonkay/code_sandbox/golang/bragdoc/file"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qustavo/dotsql"
)

type SqliteHandler struct{}

func (f SqliteHandler) Connect(config config.Config) DbHandle {
	dbHandle := connectDb(bragDbPath(
		config.Brag.Home,
		config.Brag.Dir,
		config.Sqlite.DatabaseName,
	))
	return dbHandle
}

// Extension of DbHandler struct
type SqliteDbHandle struct {
	db      *sql.DB
	version string
	dbPath  string
}

func (f SqliteDbHandle) Db() *sql.DB {
	return f.db
}

func (f SqliteDbHandle) Version() string {
	return f.version
}

func (f SqliteDbHandle) DbPath() string {
	return f.dbPath
}

func (f SqliteDbHandle) Create() {
	InitializeDb(f)
	SeedDb(f)
}

func bragDbPath(bragHome string, bragDir string, dbName string) string {
	rtn := filepath.Join(bragHome, bragDir, dbName)
	return rtn
}

func deleteDbFile(dbPath string) error {
	return os.Remove(dbPath)
}

func createDbFile(dbPath string) error {
	exists := file.CheckFileExists(dbPath)
	if exists {
		return nil
	} else {
		_, err := os.Create(dbPath)
		return err
	}
}

func connectDb(dbPath string) *SqliteDbHandle {

	// err := deleteDbFile(dbPath)
	err := createDbFile(dbPath)

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Could not open connection to database", dbPath, ", error:", err)
		os.Exit(1)
	}
	return &SqliteDbHandle{
		dbPath:  dbPath,
		version: "",
		db:      db,
	}
}

func InitializeDb(handle DbHandle) {
	dot, err := dotsql.LoadFromFile("sql/init.schema")
	if err != nil {
		log.Fatal("Could not load init.schema", err)
		handle.Db().Close()
		os.Exit(1)
	}
	_, err = dot.Exec(handle.Db(), "initialize-actions")

	if err != nil {
		log.Fatal("Could not initialize database ", err)
		handle.Db().Close()
		os.Exit(1)
	}
}

func SeedDb(handle DbHandle) {
	dot, err := dotsql.LoadFromFile("sql/init_data.sql")
	if err != nil {
		log.Fatal("Could not load init_data.sql", err)
		os.Exit(1)
	}
	_, err = dot.Exec(handle.Db(), "seed-data")

	if err != nil {
		log.Fatal("Could not Seed database", err)
		handle.Db().Close()
		os.Exit(1)
	}
}
