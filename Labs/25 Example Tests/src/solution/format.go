package format

import (
	"fmt"
	"strings"
)

// section: format

// Template takes a template and inserts values from the corresponding map
func Template(template string, values map[string]interface{}) string {
	for k, v := range values {
		template = strings.ReplaceAll(template, "%"+k, fmt.Sprintf("%v", v))
	}
	return template
}

// section: format
