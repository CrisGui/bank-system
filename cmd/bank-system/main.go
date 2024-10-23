package main

import (
	"github.com/CrisGui/bank-system/internal/config"
	"github.com/CrisGui/bank-system/internal/database"
)

func main() {
	config.RunConfig()
	database.RunDatabase()
}
