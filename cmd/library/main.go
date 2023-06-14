package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Times struct {
	N int `json:"n"`
}

func main() {
	str := os.Args
	// Read the JSON file
	fileBytes, err := ioutil.ReadFile("./times.json")
	if err != nil {
		log.Fatal(err)
	}

	// Create a variable to hold the parsed data
	var time Times

	// Unmarshal the JSON data into the variable
	err = json.Unmarshal(fileBytes, &time)
	if err != nil {
		log.Fatal(err)
	}
	if time.N == 1 {
		FileCreationsPart(str)
		time.N = 2

		// Marshal the updated data back to JSON
		updatedJSON, err := json.MarshalIndent(time, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile("./times.json", updatedJSON, 0644)
		if err != nil {
			log.Fatal(err)
		}
	} else if time.N == 2 {
		RetriveAndInsertPart(str)
	}
}
