package main

import (
	"fmt"
	"log"
	"optitech/database"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var help = `Usage:
      repository-cli <command> [arguments]

The commands are:
migrate         Run the migrations
      up		  Run the Up migrations files
	  	down      Run the Down migrations files
seed        Run the migrations
      up		  Run the Up seeders files
	  	down      Run the Down seeders files`

func main() {

	if len(os.Args) < 2 {
		log.Fatal(help)
	}

	switch os.Args[1] {
	case "migrate":
		if len(os.Args) < 3 {
			log.Printf("You must specify the argument to migrate command")
			log.Fatal(help)
		}
		migrateCommand(os.Args[2])
	case "seed":
		if len(os.Args) < 3 {
			log.Printf("You must specify the argument to seed command")
			log.Fatal(help)
		}

		err := database.Seeder(os.Args[2])

		if err != nil {
			log.Fatalf("%v", err)
		}

	default:
		log.Fatal(help)
	}

}

var m *migrate.Migrate

func migrateCommand(arg string) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting current directory")
	}

	p := fmt.Sprintf("file://%s", filepath.ToSlash(filepath.Join(wd, "database", "schemas")))

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("%v", err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		log.Fatalf("%v", err)
	}

	m, err = migrate.NewWithDatabaseInstance(
		p,
		"postgres", driver)

	if err != nil {
		log.Fatalf("%v", err)
	}

	switch arg {
	case "up":
		migrateUp()
	case "down":
		migrateDown()
	default:
		log.Fatal(help)
	}
}

func migrateUp() {
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("%v", err)
	}
	log.Print("Migration up succefully")
}

func migrateDown() {
	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("%v", err)
	}
	log.Print("Migration down succefully")
}
