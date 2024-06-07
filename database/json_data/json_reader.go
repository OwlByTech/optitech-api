package jsonreader

import (
	"encoding/json"
	"os"
)

func ReadFromJSON(fileName string, dto interface{}) error {
	jsonFile, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonFile, dto)
	if err != nil {
		return err
	}
	return nil
}
