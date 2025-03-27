package repo

import "database/sql"

type RepoUser struct {
	db *sql.DB
}

func NewRepoUser(db *sql.DB) *RepoUser {
	return &RepoUser{db}
}
