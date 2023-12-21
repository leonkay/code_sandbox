package tracker

import (
	"database/sql"

	"github.com/leonkay/code_sandbox/golang/bragdoc/context"
	"github.com/leonkay/code_sandbox/golang/bragdoc/model"
	"github.com/qustavo/dotsql"
)

type TrackerHandler interface {
	New(action model.Action, trackerType model.TrackerType, args []string) *ActionHandler
}

type StatementKeyFunc func(*context.Context) string

type ContextUpdateFunc func(recordId int, con *context.Context)

type ActionArgsFunc func(args []string, con *context.Context) []any

type ActionHandler struct {
	Action        model.Action
	SqlFile       string
	StatementKey  string
	CliArgs       []string
	ChangeDb      bool
	ContextUpdate ContextUpdateFunc
	Statement     StatementKeyFunc
	ArgsGenerator ActionArgsFunc
}

func (handle ActionHandler) Exec(con *context.Context) (int, error) {
	dot, err := dotsql.LoadFromFile(handle.SqlFile)
	if err != nil {
		return 0, err
	}

	args := handle.ArgsGenerator(handle.CliArgs, con)
	var id int64
	if handle.ChangeDb {
		res, err := dot.Exec((*con.Handler).Db(), handle.StatementKey, args...)
		if id, err = res.LastInsertId(); err != nil {
			return 0, err
		}
	} else {
		row, err := dot.QueryRow((*con.Handler).Db(), handle.StatementKey, args...)
		if err == sql.ErrNoRows {
			return 0, err
		}
		row.Scan(&id)
	}

	if id != 0 && handle.ContextUpdate != nil {
		handle.ContextUpdate(int(id), con)
	}
	return int(id), nil
}

var trackerMap = map[string]TrackerHandler{
	"company": &CompanyTrackerHandler{},
	"title":   &TitleTrackerHandler{},
	"salary":  &SalaryTrackerHandler{},
}

func TrackerMap() map[string]TrackerHandler {
	return trackerMap
}
