package main

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir, dirErr := os.UserConfigDir()

	var (
		configPath string
		origConfig []byte
	)

	if dirErr == nil {
		configPath = filepath.Join(dir, "ExampleUserConfigDir", "example.conf")

		var err error
		origConfig, err = os.ReadFile(configPath)
		if err != nil && !os.IsNotExist(err) {
			// The user has a config file but we couldn't read it.
			// Report the error instead of ignoring their configuration.
			log.Fatal(err)
		}
	}

	// Use and perhaps make changes to the config.
	config := bytes.Clone(origConfig)

	// Save changes.
	if !bytes.Equal(config, origConfig) {
		if configPath == "" {
			log.Printf("not saving config changes: %v", dirErr)
		} else {
			err := os.MkdirAll(filepath.Dir(configPath), 0700)
			if err == nil {
				err = os.WriteFile(configPath, config, 0600)
			}

			if err != nil {
				log.Printf("error saving config changes: %v", err)
			}
		}
	}
}
