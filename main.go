package main

import (
	"context"
	"fmt"
	"log"

	"MODULE_NAME/configs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	ctx := context.Background()

	cfg := configs.NewConfig()

	db, err := sqlx.ConnectContext(ctx, "mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName))

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := app.Listen(":9000"); err != nil {
		log.Fatal(err)
	}
}