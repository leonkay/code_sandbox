package tracker

import (
	"github.com/leonkay/code_sandbox/golang/bragdoc/context"
	"github.com/leonkay/code_sandbox/golang/bragdoc/model"
)

type SalaryTrackerHandler struct {
}

func (c SalaryTrackerHandler) New(action model.Action, trackerType model.TrackerType, args []string) *ActionHandler {
  if action == model.Add {
		return &ActionHandler{
			Action:       action,
			SqlFile:      "sql/salary.sql",
			StatementKey: "insert",
			CliArgs:      args,
			ChangeDb:     true,
			ArgsGenerator: func(args []string, con *context.Context) []any {
				rtn := []any{con.Process.Title.Id}
				for _, x := range args {
					rtn = append(rtn, x)
				}
				return rtn
			},
		}
	}
  return nil
}
