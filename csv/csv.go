package csv

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
)

var PATTERN = `https?:\/\/([\w-\.]+)/.*`
var re = regexp.MustCompile(PATTERN)


func GetDomainFromUrl(url string) string {
	matches := re.FindStringSubmatch(url)
	if matches != nil {
		return matches[1]
	}
	return ""
}

// GetAllDoaminNamesFromCsv reads a CSV file at the given path and extracts domain names from a specified column.
//
// Parameters:
// - p: the path to the CSV file.
// - c: the index of the column containing the domain names.
//
// Returns:
// - []string: a slice containing all the extracted domain names.
func GetAllDoaminNamesFromCsv(p string, c int) []string {
	results := []int{0, 0}
	var collum_data []string

	file, err := os.Open(p)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the CSV data
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, row := range data {
		if row[c] == "" {
			continue
		}
		domain := GetDomainFromUrl(row[c])
		if domain != "" {
			collum_data = append(collum_data, domain)
			results[0]++
		} else {
			results[1]++
		}
	}
	fmt.Printf("Total: %d, No match: %d (%2.2f%%)\n", results[0]+results[1], results[1], 100*float64(results[0])/float64(results[1]+results[0]))
	return collum_data
}
