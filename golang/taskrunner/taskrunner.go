package main

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"

	"gopkg.in/yaml.v2"
)

// used to output additional console logging for debugging
var verbose = false

// used to block command execution of the associated task.
var silentMode = false

type Task struct {
	Name     string   `yaml:"name"`
	Commands []string `yaml:"commands"`
	Args     []string `yaml:"args"`
}

type TaskList struct {
	Tasks []Task `yaml:"tasks"`
}

/**
 * A set of default tasks that always come bundled with taskrunner
 * @returns TaskList
 */
func defaultTasks() TaskList {
	tasks := TaskList{
		Tasks: []Task{
			{Name: "go-build", Commands: []string{"go build"}, Args: []string{}},
		},
	}
	return tasks
}

/**
 * Unmarshal the tasks in the referenced file
 * @param file - the file path/name.yml that contains the tasks to be loaded
 * @returns TaskList - if the file was found and unmarshalled
 * @returns error - if an error occurred during unmarshalling
 */
func loadTasks(file string) (TaskList, error) {
	var tasks TaskList

	data, err := os.ReadFile(file)
	if err != nil {
		return tasks, err
	}

	err = yaml.Unmarshal(data, &tasks)
	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

/**
 * From the task and any arguments passed build out the actual command
 * @param task - the task object template for building out the command line command
 * @param args - an array of arguments to used when building out the command line command
 * @returns string - The generated command line command to be executed
 */
func buildTask(task Task, args ...string) []string {
	baseCmds := task.Commands

	argsStr := strings.Join(args, " ")

	rtn := []string{}

	for _, cmd := range baseCmds {
		prep := strings.ReplaceAll(cmd, "{ARGS}", argsStr)
		// replace with command line arguments
		for index, arg := range args {
			prep = strings.ReplaceAll(prep, fmt.Sprintf("{ARGS[%d]}", index), arg)
		}

		// replace with default args on the task
		for index, arg := range task.Args {
			prep = strings.ReplaceAll(prep, fmt.Sprintf("{ARGS[%d]}", index), arg)
		}
		if verbose {
			fmt.Printf("-- Command: '%s'\n", prep)
		}
		rtn = append(rtn, prep)
	}
	return rtn
}

/**
 * Executes a command line command, routing the output to std out and std err
 */
func runTask(command string) error {
	cmd := exec.Command("sh", "-c", command)
	if verbose {
		fmt.Printf("-- Running Command '%s'\n", cmd.String())
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

/**
 * From an Array of arguments (representative of the cli arguments), return the following
 * @return task - the task.name string attribute
 * @return options - the options array that were passed to the command line.
 *			Note these options must come before the task name to be associated with taskrunner
 * @return args - a string array of arguments that will be used in the command line command generation.
 */
func findTask(args ...string) (string, []string, []string) {
	task := args[0]
	rtnIndex := 1
	for index, arg := range args {
		if !strings.HasPrefix(arg, "-") {
			task = arg
			rtnIndex = index
			break
		}
	}
	return task, args[0:rtnIndex], args[rtnIndex+1:]
}

/**
 * Helper function to generate a menu
 */
func help(tasks TaskList) {
	fmt.Println("Usage: taskrunner <task>")
	fmt.Println("  <options>:")
	fmt.Println("  	-s\tSilent Mode. Prevents the task command from executing.")
	fmt.Println("  	-v\tVerbose Mode. Surfaces additional logging for debugging taskrunner.")
	fmt.Println("  <task>:")
	for _, task := range tasks.Tasks {
		fmt.Printf("    - %s\n", task.Name)
	}
}

/**
 * Main Entry point
 */
func main() {
	file := "tasks.yaml"
	defTasks := defaultTasks()
	fileTasks, err := loadTasks(file)
	if err != nil {
		fmt.Println("Error loading tasks:", err)
	}

	// combine default tasks with the tasks found in tasks.yml
	tasks := TaskList{
		Tasks: append(defTasks.Tasks, fileTasks.Tasks...),
	}

	// if not enough arguments were passed, no task was specified
	// so print help
	if len(os.Args) < 2 {
		help(tasks)
		os.Exit(1)
	}

	// find the associated task from the arguments
	taskName, options, taskArgs := findTask(os.Args[1:]...)
	if len(options) > 0 {
		silentMode = slices.Contains(options, "-s")
		verbose = slices.Contains(options, "-v")
	}
	if verbose {
		fmt.Printf("-- options: %s\n", options)
	}

	// Iterate through the task list, searching for an exact match
	// and execute it if found.
	for _, task := range tasks.Tasks {
		if task.Name == taskName {
			fmt.Printf("Running task: %s\n", task.Name)
			if verbose {
				fmt.Printf("-- with args: %s\n", taskArgs)
			}
			taskCommands := buildTask(task, taskArgs...)
			if !silentMode {
				for _, taskCommand := range taskCommands {
					err := runTask(taskCommand)
					if err != nil {
						fmt.Println("Error running task:", err)
						os.Exit(1)
					}
					return
				}
			} else {
				fmt.Println("not executed")
				os.Exit(1)
			}
		}
	}

	fmt.Printf("Task '%s' not found.\n", taskName)
	help(tasks)
	os.Exit(1)
}
