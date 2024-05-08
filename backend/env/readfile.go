package env

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(name string) (map[string]string, error) {
	projectPath := "/code"
	filePath := filepath.Join(projectPath, name)
	env := make(map[string]string)

	// Open the .env file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening %v %v", filePath, err)
	}
	defer file.Close()

	// Read lines from the file and parse key-value pairs
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)

		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			env[key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading .env file: %v", err)
	}

	return env, nil
}