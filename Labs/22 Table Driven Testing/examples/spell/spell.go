package spell

// section: function
func AutoCorrect(term string) string {
	lookup := map[string]string{
		"accaptable":   "acceptable",
		"accidentilly": "accidentally",
		"accommadate":  "accommodate",
		"acqire":       "acquire",
		"alot":         "allot",
		"amatuer":      "amateur",
		"apparint":     "apparent",
		"argumint":     "argument",
		"beleive":      "believe",
		"wierd":        "weird",
	}

	return lookup[term]
}

// section: function
