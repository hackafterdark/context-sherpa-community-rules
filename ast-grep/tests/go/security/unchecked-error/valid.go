package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func goodErrorHandling() {
	// This is good - error is checked and handled
	file, err := os.Open("test.txt")
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return
	}
	defer file.Close()

	// This is also good - error is checked inline
	if err := file.Close(); err != nil {
		log.Printf("Error closing file: %v", err)
	}

	// This is good - error is captured and checked
	data := make([]byte, 100)
	n, err := file.Read(data)
	if err != nil {
		log.Printf("Error reading file: %v", err)
		return
	}
	fmt.Printf("Read %d bytes\n", n)
}

func goodErrorAssignment() error {
	// This is good - error is returned for caller to handle
	file, err := os.Open("config.json")
	if err != nil {
		return fmt.Errorf("failed to open config: %w", err)
	}
	defer file.Close()

	var config map[string]interface{}
	data := []byte(`{"key": "value"}`)
	if err := json.Unmarshal(data, &config); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}

	return nil
}
