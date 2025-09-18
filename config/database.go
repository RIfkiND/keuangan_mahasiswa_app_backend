package config

import (   
	"github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "os"
	"keuangan/backend/internals/models"
)

var DB *gorm.DB

func ConnectGorm() error {
    // Load env vars
    err := godotenv.Load(".env")
    if err != nil {
        return err
    }

    dsn := os.Getenv("SUPABASE_DB_URL") // e.g. "postgres://user:password@host:port/dbname"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
      if err != nil {
        return err
    }
	DB = db

	err = DB.AutoMigrate(models.Models...)
    if err != nil {
        return err
    }
    DB = db
    
	return nil

}