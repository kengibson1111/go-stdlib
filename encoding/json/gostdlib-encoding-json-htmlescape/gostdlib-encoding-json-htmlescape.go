package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var out bytes.Buffer

	json.HTMLEscape(&out, []byte(`{"Name":"<b>HTML content</b>"}`))
	out.WriteTo(os.Stdout)

	fmt.Println()
}
