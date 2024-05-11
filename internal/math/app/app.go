package app

import "math-parser/internal/math/app/command"

type Application struct {
	Commands Commands
	Queries Queries
}

type Commands struct {
	Parse command.ParseHandler
}

type Queries struct {
	
}