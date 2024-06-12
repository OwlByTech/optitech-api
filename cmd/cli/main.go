package main

import (
	"fmt"
	"log"
	"optitech/database"
	"os"
	"os/exec"
	"strings"
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
convert-mjml    Convert all MJML to HTML`

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
		err := convertAllMailingTemplates()
		if err != nil {
			log.Fatalf("Error converting all MJML to HTML: %v", err)
		}
	default:
		log.Fatal(help)
	}
}

func convertAllMailingTemplates() error {
	path := "./internal/service/mailing/templates"
	items, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("error read path templates")
	}

	for _, item := range items {
		filename := item.Name()
		split := strings.Split(filename, ".")
		if len(split) != 2 {
			continue
		}

		if split[1] != "mjml" {
			continue
		}

		inputPath := fmt.Sprintf("%s/%s", path, filename)
		outputPath := fmt.Sprintf("%s/%s.html", path, split[0])
		err := convertMJMLToHTML(inputPath, outputPath)

		if err != nil {
			return err
		}
	}

	return nil
}

func convertMJMLToHTML(inputPath string, outputPath string) error {
	cmd := exec.Command("mjml", "-r", inputPath, "-o", outputPath)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error converting MJML to HTML: %w", err)
	}
	return nil
}
