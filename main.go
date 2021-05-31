package main

import (
	"travel/fxmodule"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func main() {
	godotenv.Load()
	fx.New(fxmodule.Module).Run()
}
