package repo

import (
	"database/sql"
	"log"

	"github.com/agrotention/service-user/cfg"
	_ "github.com/go-sql-driver/mysql"
)

type RepoUser struct {
	db *sql.DB
}

func NewRepoUser(db *sql.DB) *RepoUser {
	return &RepoUser{db}
}

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", cfg.DSN)
	if err != nil {
		log.Fatalf("open database error: %v", err)
	}
	return db
}
