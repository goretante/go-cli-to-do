/*
Copyright © 2025 Ante Goreta agoret00@fesb.hr
*/
package main

import (
	"github.com/goretante/go-cli-to-do/cmd"
	"github.com/goretante/go-cli-to-do/internal/db"
)

func main() {
	db.InitDB()
	cmd.Execute()
}
