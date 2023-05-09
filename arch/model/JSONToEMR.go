/*

Author : Pranav
Created on : 24-03-2023
IDE : GoLand

*/

package model

import (
	"encoding/json"
	"log"
	"os"
)

func JSONToEMR(location string) *EMR {

	// Opening JSON File
	file, err := os.Open(location)
	if err != nil {
		log.Fatalln("Error opening file:", err)
	}

	//Closing JSON file
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln("Error closing file:", err)

		}
	}(file)

	// Decoding the JSON file and storing into interface
	var data interface{}
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		log.Fatalln("Error decoding file, error is :", err)

	}

	// Encode the JSON data in pretty format
	prettyJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalln("Error encoding JSON , error is :", err)

	}

	var patientData EMR
	if err := json.Unmarshal(prettyJSON, &patientData); err != nil {
		// handle error
		log.Fatalln("Error converting interface to Object , error is :", err)

	}

	return &patientData
}
