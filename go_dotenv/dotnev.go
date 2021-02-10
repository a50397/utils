package dotnev

import (
	"bufio"
	"log"
	"os"
	"path"
	"strings"
)

// DefaultConf key, value struct
type DefaultConf struct {
	Key   string
	Value string
}

// LoadDotenv load .env file in the current directory and if values are not defined, use defaults
func LoadDotenv(defaults ...DefaultConf) bool {

	// Get CWD
	dpath, err := os.Getwd()
	ferr := err
	if err != nil {
		log.Printf("Cannot get CWD")
	} else {
		dpath = path.Join(dpath, ".env")

		file, ferr := os.Open(dpath)
		defer file.Close()

		// Read the file .env if it can be opened
		if ferr == nil {
			dfile := bufio.NewReader(file)

			for {
				line, err := dfile.ReadString('\n')
				if err != nil {
					break
				}

				config := strings.Split(line, "=")

				if len(config) == 2 {
					os.Setenv(strings.ToUpper(strings.TrimSpace(config[0])), strings.TrimSpace(config[1]))
				}
			}
		}
	}

	for _, config := range defaults {
		if config.Key != "" && config.Value != "" {
			key := strings.ToUpper(strings.TrimSpace(config.Key))
			if os.Getenv(key) == "" {
				os.Setenv(key, strings.TrimSpace(config.Value))
			}
		}
	}

	return ferr != nil
}
