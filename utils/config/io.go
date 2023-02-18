package config

import (
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"io"
	"os"
)

var Config configStruct

func Init(filepath string) {
	err := read(filepath)
	if err != nil {
		fmt.Println(err)
		fmt.Println(os.ErrNotExist)
		err = write(filepath)
		if err != nil {
			panic("Error while creating config: " + err.Error())
		}
		panic("Config file created, please fill in the blanks and restart the program.")
	}
}

func read(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	configRaw, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	err = toml.Unmarshal(configRaw, &Config)
	if err != nil {
		return err
	}
	return nil
}

func write(filepath string) error {
	var configRaw []byte
	configRaw, err := toml.Marshal(Config)
	if err != nil {
		return err
	}
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(configRaw)
	if err != nil {
		return err
	}
	return nil
}
