package service

import (
	"context"
	"fmt"
	"log"
	"math-parser/config"
	"math-parser/db"
	"math-parser/internal/math/adapters"
	"math-parser/internal/math/app"
	"math-parser/internal/math/app/command"
	"math-parser/internal/math/domain/formula"
	"math-parser/internal/math/domain/parse"
	"math-parser/internal/math/domain/visitor"

	"github.com/jackc/pgx/v5"
)

const fmtDBUrl = "postgres://%s:%s@%s:%d/%s?sslmode=%s"

func NewApplication(ctx context.Context, c *config.Conf) (app.Application, func()) {
	dbString := fmt.Sprintf(fmtDBUrl, c.DB.Username, c.DB.Password, c.DB.Host, c.DB.Port, c.DB.DBName, false)
	conn, err := pgx.Connect(ctx, dbString)
	if err != nil {
		log.Fatal(err)
	}

	queries := db.New(conn)
	repository := adapters.NewRepository(queries)
	parser := parse.NewEquationParser(&visitor.FormulaVisitorImpl{Depth: 0})
	assignParser := parse.NewAssignParser(&visitor.AssignVisitorImpl{})
	equationMemory := formula.NewEquationMemory(repository, parser)
	err = equationMemory.Load(ctx)
	if err != nil {
		log.Fatal(err)
	}
	
	return app.Application{
			Commands: app.Commands{
				Parse:             command.NewParseHandler(repository, equationMemory, parser),
				SpreadRandomField: command.NewSpreadRandomFieldHandler(repository, equationMemory, assignParser),
			},
			Queries: app.Queries{},
		}, func() {
			_ = conn.Close(ctx)
		}
}
