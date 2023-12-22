package tracker

import (
	"github.com/leonkay/code_sandbox/golang/bragdoc/context"
	"github.com/leonkay/code_sandbox/golang/bragdoc/model"
)

type TitleTrackerHandler struct {
}

func (c TitleTrackerHandler) New(action model.Action, trackerType model.TrackerType, args []string) []ActionTask {
	if action == model.Set {
		return []ActionTask{
			CreateActionTask{
				Action:       action,
				SqlFile:      "sql/title.sql",
				CreateSqlKey: "insert:title",
				CliArgs:      args,
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
			},
		}
	} else if action == model.Switch {
		return []ActionTask{
			QueryActionTask{
				Action:      action,
				SqlFile:     "sql/search.sql",
				QuerySqlKey: "title:byname:selectid",
				CliArgs:     args,
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
			},
		}
	} else {
		return nil
	}
}
