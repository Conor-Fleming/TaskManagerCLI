package main

import (
	"log"
	"path/filepath"

	"github.com/Conor-Fleming/TaskManagerCLI/cmd"
	"github.com/Conor-Fleming/TaskManagerCLI/database"
	"github.com/mitchellh/go-homedir"
)

func main() {
	dir, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}
	dbPath := filepath.Join(dir, "tasks.db")
	err = database.Init(dbPath)
	if err != nil {
		panic(err)
	}
	cmd.Execute()
}
