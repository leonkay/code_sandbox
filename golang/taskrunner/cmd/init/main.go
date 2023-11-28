package main

import (
	"fmt"
	"runtime"
	"os"
	"path/filepath"
)

func homeDefault() string {
	pwd, err := os.Getwd()
	if err != nil {
			panic(err)
	}
	exPath := filepath.Join(pwd, ".taskrunner")
	return exPath
}

func main() {

	taskRunnerHomeVar := "TASKRUNNER_HOME"

	// Note the Printf are prefixed with '#' since this is assumed to be
	// piped into a bash profile sh file of some sorte
	fmt.Printf("# Operating System: %s\n", runtime.GOOS)
	fmt.Printf("# Architecture: %s\n", runtime.GOARCH)

	// Retrieve an environment variable
	taskRunnerHome := os.Getenv(taskRunnerHomeVar)

	// Check if the environment variable is set
	if taskRunnerHome == "" {
		//fmt.Printf("'%s' is not set.\n", taskRunnerHomeVar)
		taskRunnerHome = homeDefault()
		fmt.Printf("export %s=%s\n", taskRunnerHomeVar, taskRunnerHome)

	} else {
		//fmt.Printf("%s: %s\n", taskRunnerHomeVar, taskRunnerHome)
	}
	//fmt.Println(os.Environ())
}
