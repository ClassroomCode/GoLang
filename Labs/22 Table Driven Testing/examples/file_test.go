package image_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestImageResize(t *testing.T) {
	// Fetch the contents of the fixtures directory. This acts as the "table" in
	// our table-driven test
	resizeTests, err := ioutil.ReadDir("./testdata/fixtures")
	if err != nil {
		t.Fatal("Unable to open fixture directory: err:", err)
	}

	for _, test := range resizeTests {
		// use the name of the image for the subtest name:
		t.Run(test.Name(), func(t *testing.T) {
			expected, err := os.Open(filepath.Join("./testdata/golden", test.Name()))
			if err != nil {
				t.Fatal("Unable to find matching golden file for", test.Name(), "err:", err)
			}

			actual, err := resize.Resize(filepath.Join("./testdata/fixtures", test.Name()))
			if err != nil {
				t.Fatal("Err while resizing: err:", err)
			}

			if !bytes.Equal(expected, actual) {
				t.Fatal("Resized image differs from expectation")
			}
		})
	}
}
