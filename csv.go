package utility

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
)

// separator may be '\t'
func ReadCsvFile(filePath string, separator rune) (mp [][]string, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read input file %s: %w", filePath, err)
	}
	defer func() {
		if errFile := f.Close(); errFile != nil {
			// Go 1.20+: joins parse error (if any) with close error
			err = errors.Join(err, fmt.Errorf("close %s: %w", filePath, errFile))
		}
	}()

	csvReader := csv.NewReader(f)
	if separator != rune(0) {
		csvReader.Comma = separator
	}
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("unable to parse file as CSV for %s: %w", filePath, err)
	}

	return records, nil
}

func ReadTextStringArray(filePath string) (mp []string, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read input file %s: %w", filePath, err)
	}
	defer func() {
		if errFile := f.Close(); errFile != nil {
			// Go 1.20+: joins parse error (if any) with close error
			err = errors.Join(err, fmt.Errorf("close %s: %w", filePath, errFile))
		}
	}()

	arr := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("unable to parse file TXT for %s: %w", filePath, err)
	}
	return arr, nil
}

func ReadTextStringArrayReader(file io.Reader) (mp []string, err error) {
	arr := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("unable to parse file %w", err)
	}
	return arr, nil
}
