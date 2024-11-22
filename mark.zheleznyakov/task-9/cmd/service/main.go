package main

import (
  "github.com/mrqiz/task-9/internal/database"
  "github.com/mrqiz/task-9/internal/models"
)

func main() {
  database.Connect()
  database.DB.AutoMigrate(&models.Contact{})
  database.DB.Create(&models.Contact{ Name: "Octo8", Phone: "8 (888) 888-88-88" })
}
