package src_test

import (
	"io/ioutil"
	"log"
	"testing"
)

// section: verbose
func TestVerbose(t *testing.T) {
	if testing.Verbose() {
		t.Log("put extra logging here that normally we don't care about")
	} else {
		// silence my normal loggers
		log.SetOutput(ioutil.Discard)
	}
}

// section: verbose

// section: short
func Test_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test due to short testing detected")
	}
	t.Log("Running really long integration test...")
}

// section: short

// section: parallel
func Test_Parallel(t *testing.T) {
	t.Parallel()
	t.Log("Running parallel test...")
}

// section: parallel
