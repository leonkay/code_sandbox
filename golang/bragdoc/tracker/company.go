package tracker

import (
	"strings"

	"github.com/leonkay/code_sandbox/golang/bragdoc/context"
	"github.com/leonkay/code_sandbox/golang/bragdoc/model"
)

type CompanyTrackerHandler struct {
}

func (c CompanyTrackerHandler) New(action model.Action, trackerType model.TrackerType, args []string) []ActionTask {
	if action == model.Set {
		return []ActionTask{
			CreateActionTask{
				Action:       action,
				SqlFile:      "sql/company.sql",
        QuerySqlKey: "select:company",
				CreateSqlKey: "insert:company",
				CliArgs:      args,
				ContextUpdate: func(recordId int, con *context.Context) {
					con.Brag.CompanyId = recordId
					con.UpdateFile()
				},
				ArgsGenerator: func(args []string, con *context.Context) []any {
					return []any{strings.Join(args, " ")}
				},
			},
		}
	} else if action == model.Switch {
		return []ActionTask{
			QueryActionTask{
				Action:      action,
				SqlFile:     "sql/company.sql",
				QuerySqlKey: "select:company",
				CliArgs:     args,
				ContextUpdate: func(recordId int, con *context.Context) {
					con.Brag.CompanyId = recordId
					con.UpdateFile()
				},
				ArgsGenerator: func(args []string, con *context.Context) []any {
					return []any{strings.Join(args, " ")}
				},
			},
		}
	} else if action == model.Clear {
		return []ActionTask{
			QueryActionTask{
				Action:      action,
				SqlFile:     "sql/company.sql",
				QuerySqlKey: "select:company",
				CliArgs:     args,
				ContextUpdate: func(recordId int, con *context.Context) {
					con.Brag.CompanyId = recordId
					con.UpdateFile()
				},
				ArgsGenerator: func(args []string, con *context.Context) []any {
					return []any{strings.Join(args, " ")}
				},
			},
		}
	} else {
		return nil
	}
}
