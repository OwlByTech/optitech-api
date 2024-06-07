package mailing

import (
	"fmt"
	"os/exec"
)

func ConvertMJML(mjmlFile string) error {
	outputFile := "output.html"
	cmd := exec.Command("mjml", mjmlFile, "-o", outputFile)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error converting MJML to HTML: %w", err)
	}

	fmt.Printf("Successfully converted %s to %s\n", mjmlFile, outputFile)
	return nil
}
