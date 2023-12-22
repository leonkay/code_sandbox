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

type QueryActionTask struct {
	Action        model.Action
	SqlFile       string
	QuerySqlKey   string
	CliArgs       []string
	ContextUpdate ContextUpdateFunc
	ArgsGenerator ActionArgsFunc
}

func (handle QueryActionTask) Exec(con *context.Context) (int, error) {
	dot, err := dotsql.LoadFromFile(handle.SqlFile)
	if err != nil {
		return 0, err
	}

	args := handle.ArgsGenerator(handle.CliArgs, con)
	var id int64
	row, err := dot.QueryRow((*con.Handler).Db(), handle.QuerySqlKey, args...)
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

type CreateActionTask struct {
	Action        model.Action
	SqlFile       string
	CreateSqlKey  string
	QuerySqlKey   string
	CliArgs       []string
	ContextUpdate ContextUpdateFunc
	ArgsGenerator ActionArgsFunc
}

func (handle CreateActionTask) Exec(con *context.Context) (int, error) {
	dot, err := dotsql.LoadFromFile(handle.SqlFile)
	if err != nil {
		return 0, err
	}

	var id int64
	args := handle.ArgsGenerator(handle.CliArgs, con)
	if handle.QuerySqlKey != "" {
		row, err := dot.QueryRow((*con.Handler).Db(), handle.QuerySqlKey, args...)
		if err != nil {
			if err != sql.ErrNoRows {
				return 0, err
			}
      // else tracker entity already exists so don't create a new row
      // and use the existing row
		} else {
			row.Scan(&id)
		}
	}
	if id == 0 {
		res, err := dot.Exec((*con.Handler).Db(), handle.CreateSqlKey, args...)
		if id, err = res.LastInsertId(); err != nil {
			return 0, err
		}
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
