package main

import (
	"flag"
	"log"
	"os"

	"github.com/leonkay/code_sandbox/golang/bragdoc/config"
	"github.com/leonkay/code_sandbox/golang/bragdoc/context"
	"github.com/leonkay/code_sandbox/golang/bragdoc/db"
	"github.com/leonkay/code_sandbox/golang/bragdoc/engine"
	"github.com/leonkay/code_sandbox/golang/bragdoc/file"
	"github.com/leonkay/code_sandbox/golang/bragdoc/model"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	conf := config.New()
	_, err = file.BragDir(conf)
	if err != nil {
		log.Fatal("Could not find brag doc home")
		os.Exit(1)
	}

	request := model.RequestContext{}
	flag.StringVar(&request.Company, "company", "", "The Name of the Company the Activity is for")
	flag.StringVar(&request.Project, "project", "", "The project")
	flag.StringVar(&request.Title, "title", "", "The Title this activity was for")
	flag.StringVar(&request.Role, "role", "", "The Role this activity was for")
	flag.StringVar(&request.DateStr, "date", "", "The date this activity occurred on (YYMMDD)")
	flag.Parse()
	args := flag.Args()

	if len(args) < 2 {
		log.Fatal("Please specify an action and activity as arguments")
		os.Exit(1)
	}
	action := args[0]
	trackerKey := args[1]

	request.Action = action
	request.TrackerKey = trackerKey
	request.Args = args[2:]

	handler := db.Run(*conf)

	cont := &context.Context{
		Request: &request,
		Config:  conf,
		Handler: &handler,
	}
	cont.Init()

	eng := engine.Engine{
		Context: cont,
	}
	eng.Exec()

	db.Close(handler)
	cont.UpdateFile()

	os.Exit(1)

	// p := tea.NewProgram(cli.InitialBragModel())
	//   if _, err := p.Run(); err != nil {
	// 		log.Fatalf("Alas, there's been an error: %v", err)
	// 		db.Close(handler)
	// 		os.Exit(1)
	//   }
}
