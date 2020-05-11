// +build integration

package exercise_test

import (
	"testing"
	"time"
)

func TestIntegration(t *testing.T) {
	// Long running integration test...
	time.Sleep(3 * time.Second)
	t.Log("running long integration test...")
}
