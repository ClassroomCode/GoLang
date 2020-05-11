package many

type Option int

const (
	All = iota
	First
	Last
)

// Find returns the index of the value searched for
func Find(s string, values []string, option Option) []int {
	switch option {
	case First:
		for i := range values {
			if values[i] == s {
				return []int{i}
			}
		}
		return []int{}
	case Last:
		for i := len(values) - 1; i >= 0; i-- {
			if values[i] == s {
				return []int{i}
			}
		}
		return []int{}
	default:
		r := []int{}
		for i := range values {
			if values[i] == s {
				r = append(r, i)
			}
		}
		return r
	}
}
