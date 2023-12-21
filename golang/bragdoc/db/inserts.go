package db

import (
	"github.com/leonkay/code_sandbox/golang/bragdoc/model"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qustavo/dotsql"
)

func InsertActivityEvent(handle DbHandle, action model.Action, trackerTypeId int, title string) (int, error) {
	dot, err := dotsql.LoadFromFile("sql/search.sql")
	if err != nil {
		return 0, err
	}

	res, err := dot.Exec(handle.Db(), "insert:activityevent", trackerTypeId, action, title)
	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}
	return int(id), nil
}
