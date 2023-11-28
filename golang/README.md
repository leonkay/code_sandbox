# Readme

A Taskrunner implemented in `golang`.

## Goals
Passion project to learn about `golang` with the eventual goal of building other cli tools that could be useful.

## Description

Tasks are short cuts for command line code similar to an `alias`, but allows
these tasks to be defined at the directory level.

It looks for a `tasks.yml` file in the current directory, which contains a
`TaskList` which is composed of `Tasks`. `Tasks` have the following attributes:
* name: The name of the task when running from the command line
* commands: An array of command line commands that will be executed sequentially.
  * supports pass through argument substitution using `{ARGS}` for all cli args or `{ARGS[0..n]}` for individual arguments
* args: An array of default arguments if `{ARGS|ARGS[0..n]}` are not passed.

See [tasks.yml](taskrunner/tasks.yml) for a sample. This sample has tasks specific to task runner
but if properly aliased, tasks.yml can contain tasks specific to your project.

## Requirements
This project is build in `golang@1.21.4`.

## Build
To build the project, cd into the directory e.g. `/taskrunner` and run the following command:

    $ go run taskrunner.go build-taskrunner

Alternatively the raw command can be executed

    $ go build -o dist/taskrunner

This will create the artifact `dist/taskrunner`.

## Testing
The default `tasks.yml` file contains some testing commands. In particular `echo` is used to test that `{ARGS}` are passed correctly. To test:

    $ dist/taskrunner echo This is a test
    Running task: echo
    This is a test


## Aliasing
`taskrunner` once in your path, can be used to parse a `tasks.yml` file anywhere and used to execute tasks.
