package utils

import (
	"os"
	"encoding/csv"	
	"fmt"
)

func GetCSVValueAtIndex(filename string, rowIndex int) (string, error) { 
	// unsafe
	//allows for array limit being exceeded
	
	file,err := os.Open(filename)
	if err != nil{
		fmt.Println("error opening the file!")
	}
	// Create CSV reader
	reader := csv.NewReader(file)
	
	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		return "", fmt.Errorf("failed to read CSV: %w", err)
	}
	
	// Check if file has at least a header
	if len(records) < 1 {
		return "", fmt.Errorf("CSV file is empty")
	}
	
	// Skip header (index 0), so actual data starts at records[1]
	dataRows := records[1:]
	
	// Check if the requested index is valid
	if rowIndex < 0 || rowIndex >= len(dataRows) {
		return "", fmt.Errorf("index %d out of range (0-%d)", rowIndex, len(dataRows)-1)
	}
	
	// Get the row at the specified index
	row := dataRows[rowIndex]
	
	// Check if the row has at least 3 elements
	if len(row) < 3 {
		return "", fmt.Errorf("row %d does not have a 3rd element (only has %d columns)", rowIndex, len(row))
	}
	
	return row[3], nil
}
