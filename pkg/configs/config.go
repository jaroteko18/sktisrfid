package configs

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	User     string
	Password string
	Server   string
	Database string
}

func GetConfig() Config {
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}

func (c Config) GetMyConnectionInfo() string {
	return fmt.Sprintf(
		"server=%s;user id=%s;password=%s;database=%s",
		c.Server, c.User, c.Password, c.Database)
}
