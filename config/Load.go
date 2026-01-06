package config

import (
	"encoding/json"
	"os"
)

type users struct {
	Users []struct {
		Password string `json:"password"`
		Username string `json:"username"`
		//		Interface bool   `json:"interface"`
	} `json:"users"`
}

var (
	Users *users
)

func Load() {
	loadUsers()
}

func loadUsers() error {
	f, err := os.ReadFile("./config/settings.json")
	if err != nil {
		return err
	}
	if err = json.Unmarshal(f, &Users); err != nil {
		return err
	}
	return nil
}
