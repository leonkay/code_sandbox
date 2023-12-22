package tracker

import (
	"log"

	"github.com/leonkay/code_sandbox/golang/bragdoc/context"
	"github.com/leonkay/code_sandbox/golang/bragdoc/model"
)

type SalaryTrackerHandler struct {
}

func salaryArgGenerator(args []string, con *context.Context) []any {
	log.Println(args, con)
	rtn := []any{}
	if con.Process.Title != nil {
		rtn = append(rtn, con.Process.Title.Id)
	} else {
		rtn = append(rtn, "")
	}
	for _, x := range args {
		rtn = append(rtn, x)
	}
	return rtn
}

func (c SalaryTrackerHandler) New(action model.Action, trackerType model.TrackerType, args []string) []ActionTask {
	if action == model.Set {
		return []ActionTask{
			CreateActionTask{
				Action:        action,
				SqlFile:       "sql/salary.sql",
				CreateSqlKey:  "insert",
				CliArgs:       args,
				ArgsGenerator: salaryArgGenerator,
			},
		}
	}
	return nil
}
