package main

import (
	"fmt"
	"regexp"
)

func main() {
	content := `
	# comment line
	option1: value1
	option2: value2

	# another comment line
	option3: value3
`

	// Regex pattern captures "key: value" pair from the content.
	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)

	// Template to convert "key: value" to "key=value" by
	// referencing the values captured by the regex pattern.
	template := "$key=$value\n"

	result := []byte{}

	// For each match of the regex in the content.
	for _, submatches := range pattern.FindAllStringSubmatchIndex(content, -1) {
		// Apply the captured submatches to the template and append the output
		// to the result.
		result = pattern.ExpandString(result, template, content, submatches)
	}

	fmt.Println(string(result))
}
