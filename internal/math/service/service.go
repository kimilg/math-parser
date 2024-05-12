package service

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"math-parser/config"
	"math-parser/db"
	"math-parser/internal/math/adapters"
	"math-parser/internal/math/app"
	"math-parser/internal/math/app/command"
	"math-parser/internal/math/domain/formula"
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
	equationMemory := formula.NewEquationMemory(repository)
	err = equationMemory.Load(ctx)
	if err != nil {
		log.Fatal(err)
	}
	
	return app.Application{
		Commands: app.Commands{
			Parse: command.NewParseHandler(repository),
			SpreadRandomField: command.NewSpreadRandomFieldHandler(repository, equationMemory),
		},
		Queries: app.Queries{
			
		},
	}, func() {
			_ = conn.Close(ctx)	
		}
}