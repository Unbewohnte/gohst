/*
            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
                    Version 2, December 2004

 Copyright (C) 2004 Sam Hocevar <sam@hocevar.net>

 Everyone is permitted to copy and distribute verbatim or modified
 copies of this license document, and changing it is allowed as long
 as the name is changed.

            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION

  0. You just DO WHAT THE FUCK YOU WANT TO.
*/

// Kasyanov N.A. (Unbewohnte), 2023

package conf

import (
	"encoding/json"
	"io"
	"os"
)

type Conf struct {
	Port           uint16 `json:"port"`
	CertFilePath   string `json:"cert_file_path"`
	KeyFilePath    string `json:"key_file_path"`
	BaseContentDir string `json:"base_content_dir"`
	ProdDBName     string `json:"production_db"`
}

// Creates a default server configuration
func Default() Conf {
	return Conf{
		Port:           8080,
		CertFilePath:   "",
		KeyFilePath:    "",
		BaseContentDir: ".",
		ProdDBName:     "database.db",
	}
}

// Tries to retrieve configuration from given json file
func FromFile(path string) (Conf, error) {
	configFile, err := os.Open(path)
	if err != nil {
		return Default(), err
	}
	defer configFile.Close()

	confBytes, err := io.ReadAll(configFile)
	if err != nil {
		return Default(), err
	}

	var config Conf
	err = json.Unmarshal(confBytes, &config)
	if err != nil {
		return Default(), err
	}

	return config, nil
}

// Create empty configuration file
func Create(path string, conf Conf) (Conf, error) {
	configFile, err := os.Create(path)
	if err != nil {
		return Default(), err
	}
	defer configFile.Close()

	configJsonBytes, err := json.MarshalIndent(conf, "", " ")
	if err != nil {
		return conf, err
	}

	_, err = configFile.Write(configJsonBytes)
	if err != nil {
		return conf, nil
	}

	return conf, nil
}
