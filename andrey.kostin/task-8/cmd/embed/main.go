package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed data.json
var jsonData []byte

type Preferences struct {
	Notifications bool   `json:"notifications"`
	Language      string `json:"language"`
}

type UserProfile struct {
	Username    string      `json:"username"`
	Age         int         `json:"age"`
	Email       string      `json:"email"`
	Preferences Preferences `json:"preferences"`
}

func main() {
	var profile UserProfile

	err := json.Unmarshal(jsonData, &profile)
	if err != nil {
		fmt.Println("Ошибка чтения JSON:", err)
		return
	}

	fmt.Printf("Профиль пользователя:\n")
	fmt.Printf("Имя пользователя: %s\n", profile.Username)
	fmt.Printf("Возраст: %d\n", profile.Age)
	fmt.Printf("Email: %s\n", profile.Email)
	fmt.Printf("Предпочтения:\n")
	fmt.Printf("  Уведомления: %t\n", profile.Preferences.Notifications)
	fmt.Printf("  Язык: %s\n", profile.Preferences.Language)
}
