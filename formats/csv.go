package formats

import (
	"encoding/csv"
	"os"
)

type CSVReader struct {
	File     *os.File
	Reader   *csv.Reader
	Headings []string
}

func ReadCSVFile(path string, separator rune) (reader *CSVReader, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	reader = new(CSVReader)
	reader.Reader = csv.NewReader(file)
	reader.Reader.Comma = separator
	reader.Reader.LazyQuotes = true

	return reader, nil
}

func (reader *CSVReader) Read() (record map[string]string, err error) {
	r, err := reader.Reader.Read()
	if err != nil {
		return nil, err
	}

	record = make(map[string]string)
	for i := 0; i < len(r); i++ {
		if i < len(reader.Headings) {
			record[reader.Headings[i]] = r[i]
		} else {
			record[string(i)] = r[i]
		}
	}

	return record, nil
}
