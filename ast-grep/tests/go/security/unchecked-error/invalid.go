package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func badErrorHandling() {
	// This is bad - error is ignored
	file, err := os.Open("test.txt")
	if err != nil {
		// Error is logged but execution continues without proper handling
		fmt.Printf("Error: %v\n", err)
	}
	defer file.Close()

	// This is bad - error is completely ignored
	file.Close()

	// This is bad - error is assigned but never checked
	data := make([]byte, 100)
	n, err := file.Read(data)
	fmt.Printf("Read %d bytes\n", n)
	// err is never checked!

	// This is bad - multiple errors are ignored
	os.Remove("nonexistent.txt")
	os.Mkdir("test", 0755)
	json.Unmarshal([]byte("invalid"), nil)
}

func badErrorAssignment() {
	// This is bad - error is captured but never checked
	file, err := os.Open("config.json")
	defer file.Close()

	// Error is never checked!
	var config map[string]interface{}
	data := []byte(`{"key": "value"}`)
	json.Unmarshal(data, &config)

	// This is bad - error is ignored in assignment
	result, err := someFunction()
	fmt.Printf("Result: %v\n", result)
	// err is never checked!

	// This is bad - multiple unchecked errors
	file2, _ := os.Open("another.txt") // Error ignored with _
	file2.Close()
	os.Remove("temp.txt")
}

func someFunction() (string, error) {
	return "result", nil
}
