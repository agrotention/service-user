package main

import (
	"flag"
	"log"

	"github.com/agrotention/service-user/cfg"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	isDown := flag.Bool("down", false, "use -down to revert migration")
	flag.Parse()

	m, err := migrate.New("file://migrations", cfg.DSN)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	defer m.Close()

	if *isDown {
		log.Println("Migration down...")
		if err := m.Down(); err != nil {
			log.Fatalf("Migration down failed: %v", err)
		}
	} else {
		log.Println("Migration up...")
		if err := m.Up(); err != nil {
			log.Fatalf("Migration up failed: %v", err)
		}
	}

	log.Println("Migration successful!")
}
