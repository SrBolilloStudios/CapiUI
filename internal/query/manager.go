package query

import (
	"embed"
	"fmt"
	"strings"
)

//go:embed *.sql
var sqlFiles embed.FS

var store map[string]string

func Load() error {
	store = make(map[string]string)
	files, err := sqlFiles.ReadDir(".")
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		content, err := sqlFiles.ReadFile(file.Name())
		if err != nil {
			return err
		}
		parseFile(string(content))
	}
	return nil
}

func parseFile(content string) {
	lines := strings.Split(content, "\n")
	var name string
	var statement []string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		switch {
		case strings.HasPrefix(line, "--name:"):
			if name != "" {
				store[name] = strings.Join(statement, " ")
			}
			name = strings.TrimSpace(strings.TrimPrefix(line, "--name:"))
			statement = nil
		case line != "" && !strings.HasPrefix(line, "--"):
			statement = append(statement, line)
		}
	}
	if name != "" && len(statement) > 0 {
		store[name] = strings.Join(statement, " ")
	}
}

func Get(name string) (string, error) {
	q, ok := store[name]
	if !ok {
		return "", fmt.Errorf("query %q not found", name)
	}
	return q, nil
}
