package main

import (
	"fmt"
	"path/filepath"
	"os"
	"testing"
)

func TestReadShouldReadGClogAndParse(t *testing.T) {
	var curDir, _ = os.Getwd()
	path := filepath.Join(curDir, "fixtures", "gc.log.0")
	read(path)
	for k := range logs {
		if k != "2015-3-23" && k != "2015-3-24" {
			t.Fatalf("Expected %v, but %v:", "2015-3-23 or 2015-3-24", k)
		}
	}

	logs = map[string][]string{}
}

func TestParseShouldParseLine(t *testing.T) {
	testString := "2015-03-25T23:48:54.474+0900: 5.700:"
	parse(testString)
	for k := range logs {
		if k != "2015-3-25" {
			t.Fatalf("Expected %v, but %v:", "2015-03-25", k)
		}
	}
	logs = map[string][]string{}
}

func TestWriteShouldGenerateNewFile(t * testing.T) {
	var curDir, _ = os.Getwd()
	path := filepath.Join(curDir, "fixtures", "gc.log.0")
	read(path)
	dir := filepath.Dir(path)
	fileName := filepath.Base(path)
	for k := range logs {
		write(dir, fmt.Sprintf("%s_%s", k, fileName), logs[k])
	}

	expects := []string{
		filepath.Join(curDir, "fixtures", "2015-3-23_gc.log.0"),
		filepath.Join(curDir, "fixtures", "2015-3-24_gc.log.0"),
	}

	for i := range expects {
		_, err := os.Stat(expects[i])
		if err != nil {
			t.Fatalf("Expected %v, but not existed.", expects[i])
		}
		if err := os.Remove(expects[i]); err != nil {
			t.Fatalf("Remove failed.")
		}
	}

	logs = map[string][]string{}
}
