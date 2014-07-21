package irtools

import (
	"bufio"
	"os"
	"strings"
)

// Take a list of terms and produce a frequency map for the list
func Count(terms []string) (counted map[string]int) {
	counted = make(map[string]int)
	for i := range terms {
		counted[terms[i]]++
	}
	return counted
}

// Take two term frequency maps add the source map to the destination map
func Combine(dest, source *map[string]int) {
	d := *dest
	for k, v := range *source {
		d[k] += v
	}
}

// Counts the number of terms in m2 which are not in m1
func Difference(m1, m2 map[string]int) (count int, difference []string) {
	for k, _ := range m2 {
		if _, ok := m1[k]; ok != true {
			difference = append(difference, k)
		}
	}
	return len(difference), difference
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func URLFilter(text string) (clean string, urls []string) {
	clean = text
	for {
		start := strings.Index(clean, "http://")
		if start == -1 {
			start = strings.Index(clean, "https://")
		}
		if start >= 0 {
			end := strings.Index(clean[start:], " ")
			if end == -1 {
				urls = append(urls, clean[start:])
				clean = clean[:start]
				break
			} else {
				end += start + 1
				urls = append(urls, clean[start:end])
				clean = clean[:start] + clean[end:]
			}
		} else {
			break
		}
	}
	return clean, urls
}
