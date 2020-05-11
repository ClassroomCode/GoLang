package main

// section: scribe
type scribe struct {
	data []byte
}

// implements the io.Writer interface
func (s *scribe) Write(p []byte) (int, error) {
	s.data = p
	return len(p), nil
}

// implements the fmt.Stringer interface
func (s scribe) String() string {
	return string(s.data)
}

// section: scribe
