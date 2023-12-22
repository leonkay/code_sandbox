package engine

import (
	"errors"
	"log"
	"strings"

	"github.com/leonkay/code_sandbox/golang/bragdoc/context"
	"github.com/leonkay/code_sandbox/golang/bragdoc/db"
	"github.com/leonkay/code_sandbox/golang/bragdoc/model"
	"github.com/leonkay/code_sandbox/golang/bragdoc/tracker"
)

type Process struct {
	Action  model.Action
	Tracker model.TrackerType
	Args    []string
}

func (p Process) Validate(context *context.Context) error {
	log.Println("ProcessContext: ", context.Process)
	log.Println("\tprocess: ", p)
  companyTask :=  p.Tracker.Key == "company"
	if context.Process.Company == nil && !companyTask {
		return errors.New("Please run `-switch=true set company <company name>' first")
	}
	if !companyTask && context.Process.Title == nil && p.Tracker.Key != "title" {
		return errors.New("Please run `-switch=true set title <title> <level>' first")
	}
	return nil
}

func (p Process) Run(context *context.Context) error {
	trackerHandler := tracker.TrackerMap()[p.Tracker.Key]
	if trackerHandler != nil {
		actionHandlers := trackerHandler.New(p.Action, p.Tracker, p.Args)
		if len(actionHandlers) > 0 {
			for _, actionHandler := range actionHandlers {
				_, err := actionHandler.Exec(context)
				if err != nil {
					log.Fatalln("Could not execute")
				}
			}
		} else {
			log.Println("No Specific Action Handler for: ", p)
		}
	} else {
		log.Println("No Specific Tracker Handler for: ", p)
	}

	if !model.IsIgnoreActivityEvent(p.Action) {
		_, err := db.InsertActivityEvent(*context.Handler, p.Action, p.Tracker.Id, strings.Join(p.Args, " "))
		if err != nil {
			log.Println("Could not insert activity", err)
		}
		return err
	}
	return nil
}

type Engine struct {
	Context *context.Context
}

func (e Engine) Exec() {
	req := e.Context.Request

	args := req.Args

	actionStr := req.Action
	trackerKey := req.TrackerKey

	action, err := model.ActionString(actionStr)
	if err != nil {
		log.Fatal("Could not parse action: ", actionStr)
		return
	}

	var tracker model.TrackerType
	for _, x := range e.Context.Cache.Trackers {
		if x.Key == trackerKey {
			tracker = x
		}
	}

	process := Process{
		Action:  action,
		Args:    args,
		Tracker: tracker,
	}

	if err := process.Validate(e.Context); err == nil {
		process.Run(e.Context)
	} else {
		log.Fatalln("Process Validate Error: ", err)
	}
}
