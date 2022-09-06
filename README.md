# TaskManagerCLI

[![Go Report Card](https://goreportcard.com/badge/github.com/Conor-Fleming/TaskManagerCLI)](https://goreportcard.com/report/github.com/Conor-Fleming/TaskManagerCLI)

## About
Task Manager CLI is a tool that can help you keep track of your tasks. This tool utilizes Cobra to help with the command line interface and taking arguments. It also uses BoltDB to handle the storing of tasks.

*Project status: Working/In Progress

## Installation and Use
*Clone Github Repo: https://github.com/Conor-Fleming/TaskManagerCLI
```
git clone https://github.com/Conor-Fleming/TaskManagerCLI
```
Use go build to create an executable:
```
go build -o task
```

You can use a blank ./task call in the terminal to view help test and see the available commands:
```
./task
```
```
Task is a CLI to help you manage tasks

Usage:
  task [command]

Available Commands:
  add         Use add to add a new task to the TODO list
  clear       Remove all the tasks from the list
  completion  Generate the autocompletion script for the specified shell
  done        Use done to mark a task as complete and remove from the list
  help        Help about any command
  list        View a list of task to be completed

Flags:
  -h, --help      help for task
  -v, --version   version for task

Use "task [command] --help" for more information about a command.
```
Use any of the available commands with:
```
./task [command]
```
