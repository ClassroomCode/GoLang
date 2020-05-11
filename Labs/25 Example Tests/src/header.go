package src

import "fmt"

func Header(h string, lines []string) string {
	s := fmt.Sprintf("<h1>%s</h1>\n", h)
	for _, l := range lines {
		s = fmt.Sprintf("%s\n%s", s, l)
	}
	return s
}
