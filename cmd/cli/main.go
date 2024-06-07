package main

import (
	"log"
	"optitech/database"
	"optitech/internal/service/mailing"
	"os"
)

var help = `Usage:
      repository-cli <command> [arguments]

The commands are:
migrate         Run the migrations
      up		  Run the Up migrations files
	  	down      Run the Down migrations files
seed        Run the migrations
      up		  Run the Up seeders files
	  	down      Run the Down seeders files
convert-mjml    Convert MJML to HTML
    <file>        ./internal/service/mailing/templates`

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

		err := database.Migrate(os.Args[2])
		if err != nil {
			log.Fatalf("%v", err)
		}
	case "seed":
		if len(os.Args) < 3 {
			log.Printf("You must specify the argument to seed command")
			log.Fatal(help)
		}

		err := database.Seeder(os.Args[2])

		if err != nil {
			log.Fatalf("%v", err)
		}

	case "convert-mjml":
		if len(os.Args) < 3 {
			log.Printf("You must specify the MJML file path")
			log.Fatal(help)
		}

		err := mailing.ConvertMJML(os.Args[2])
		if err != nil {
			log.Fatalf("Error converting MJML to HTML: %v", err)
		}

	default:
		log.Fatal(help)
	}

}
