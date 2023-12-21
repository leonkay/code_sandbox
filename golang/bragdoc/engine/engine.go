package engine

import (
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

	var err error

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

	process.Run(e.Context)
}
