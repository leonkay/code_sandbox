package tracker

import (
	"strings"

	"github.com/leonkay/code_sandbox/golang/bragdoc/context"
	"github.com/leonkay/code_sandbox/golang/bragdoc/model"
)

type CompanyTrackerHandler struct {
}

func (c CompanyTrackerHandler) New(action model.Action, trackerType model.TrackerType, args []string) *ActionHandler {
	if action == model.Join {
		return &ActionHandler{
			Action:       action,
			SqlFile:      "sql/company.sql",
			StatementKey: "insert:company",
			CliArgs:      args,
			ChangeDb:     true,
			ContextUpdate: func(recordId int, con *context.Context) {
				con.Brag.CompanyId = recordId
				con.UpdateFile()
			},
			ArgsGenerator: func(args []string, con *context.Context) []any {
				return []any{strings.Join(args, " ")}
			},
		}
	} else if action == model.Add {
		return &ActionHandler{
			Action:       action,
			SqlFile:      "sql/company.sql",
			StatementKey: "insert:company",
			CliArgs:      args,
			ChangeDb:     true,
			ArgsGenerator: func(args []string, con *context.Context) []any {
				return []any{strings.Join(args, " ")}
			},
		}
	} else {
		return nil
	}
}
