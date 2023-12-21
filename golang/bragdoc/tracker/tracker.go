package tracker

import (
	"database/sql"

	"github.com/leonkay/code_sandbox/golang/bragdoc/context"
	"github.com/leonkay/code_sandbox/golang/bragdoc/model"
	"github.com/qustavo/dotsql"
)

type ActionTask interface {
  Exec(con *context.Context) (int, error)
}

type ActionTaskGenerator interface {
	New(action model.Action, trackerType model.TrackerType, args []string) []ActionTask
}
type StatementKeyFunc func(*context.Context) string

type ContextUpdateFunc func(recordId int, con *context.Context)

type ActionArgsFunc func(args []string, con *context.Context) []any

type QueryActionHandler struct {
	Action        model.Action
	SqlFile       string
	StatementKey  string
	CliArgs       []string
	ContextUpdate ContextUpdateFunc
	Statement     StatementKeyFunc
	ArgsGenerator ActionArgsFunc
}

type UpdateActionHandler struct {
	Action        model.Action
	SqlFile       string
	StatementKey  string
	CliArgs       []string
	ContextUpdate ContextUpdateFunc
	Statement     StatementKeyFunc
	ArgsGenerator ActionArgsFunc
}


func (handle QueryActionHandler) Exec(con *context.Context) (int, error) {
	dot, err := dotsql.LoadFromFile(handle.SqlFile)
	if err != nil {
		return 0, err
	}

	args := handle.ArgsGenerator(handle.CliArgs, con)
	var id int64
  row, err := dot.QueryRow((*con.Handler).Db(), handle.StatementKey, args...)
  if err == sql.ErrNoRows {
    return 0, err
  }
  row.Scan(&id)

	if con.Request.SwitchContext && handle.ContextUpdate != nil {
		if id != 0 {
			handle.ContextUpdate(int(id), con)
		}
	}

	return int(id), nil
}

func (handle UpdateActionHandler) Exec(con *context.Context) (int, error) {
	dot, err := dotsql.LoadFromFile(handle.SqlFile)
	if err != nil {
		return 0, err
	}

	args := handle.ArgsGenerator(handle.CliArgs, con)
	var id int64
  res, err := dot.Exec((*con.Handler).Db(), handle.StatementKey, args...)
  if id, err = res.LastInsertId(); err != nil {
    return 0, err
  }

	if con.Request.SwitchContext && handle.ContextUpdate != nil {
		if id != 0 {
			handle.ContextUpdate(int(id), con)
		}
	}

	return int(id), nil
}

var trackerMap = map[string]ActionTaskGenerator{
	"company": CompanyTrackerHandler{},
	"title":   TitleTrackerHandler{},
	"salary":  SalaryTrackerHandler{},
}

func TrackerMap() map[string]ActionTaskGenerator {
	return trackerMap
}
