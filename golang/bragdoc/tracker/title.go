package tracker

import (
	"github.com/leonkay/code_sandbox/golang/bragdoc/context"
	"github.com/leonkay/code_sandbox/golang/bragdoc/model"
)

type TitleTrackerHandler struct {
}

func (c TitleTrackerHandler) New(action model.Action, trackerType model.TrackerType, args []string) *ActionHandler {
	if action == model.Join {
		return &ActionHandler{
			Action:       action,
			SqlFile:      "sql/title.sql",
			StatementKey: "insert:title",
			CliArgs:      args,
			ChangeDb:     true,
			ContextUpdate: func(recordId int, con *context.Context) {
				con.Brag.TitleId = recordId
				con.UpdateFile()
			},
			ArgsGenerator: func(args []string, con *context.Context) []any {
				rtn := []any{con.Process.Company.Id}
				for _, x := range args {
					rtn = append(rtn, x)
				}
				return rtn
			},
		}
	} else if action == model.Add {
		return &ActionHandler{
			Action:       action,
			SqlFile:      "sql/title.sql",
			StatementKey: "insert:title",
			CliArgs:      args,
			ChangeDb:     true,
			ArgsGenerator: func(args []string, con *context.Context) []any {
				rtn := []any{con.Process.Company.Id}
				for _, x := range args {
					rtn = append(rtn, x)
				}
				return rtn
			},
		}
	} else if action == model.Switch {
		return &ActionHandler{
			Action:       action,
			SqlFile:      "sql/search.sql",
			StatementKey: "title:byname:selectid",
			CliArgs:      args,
			ChangeDb:     false,
			ContextUpdate: func(recordId int, con *context.Context) {
				con.Brag.TitleId = recordId
				con.UpdateFile()
			},
			ArgsGenerator: func(args []string, con *context.Context) []any {
				rtn := []any{con.Process.Company.Id}
				for _, x := range args {
					rtn = append(rtn, x)
				}
				return rtn
			},
		}
	} else {
		return nil
	}
}
