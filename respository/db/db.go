package db

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/go-pg/pg"
)

type PgDbCfg struct {
	User     string `env:"PG_DB_USER" envDefault:"postgres"`
	Address  string `env:"PG_DB_ADDRESS" envDefault:"127.0.0.1"`
	Password string `env:"PG_DB_PASSWORD"`
	Database string `env:"PG_DB_DATABASE" envDefault:"postgres"`
}

func NewPgDb() *pg.DB {
	cfg := PgDbCfg{}

	if err := env.Parse(&cfg); err != nil {
		fmt.Println("Error loading PgDb cfg.", "Error:", err)
		panic(err)
		return nil
	}
	db := pg.Connect(&pg.Options{
		Addr:     cfg.Address,
		User:     cfg.User,
		Password: cfg.Password,
		Database: cfg.Database,
	})

	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS "users" (
		"id" SERIAL PRIMARY KEY,
		"name" VARCHAR(128) NOT NULL,
    	"email" VARCHAR(128) NOT NULL,
    	"password" VARCHAR(255) NOT NULL,
    	"phone_no" VARCHAR(15) NOT NULL
	  );`)

	if err != nil {
		panic(err)
		return nil
	}

	fmt.Printf("Connected to Postgres DB.\n")

	return db
}
